package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

const (
	// bufferTime 缓冲时间，1天内自动续期
	bufferTime = 86400
)

// JwtClaims JWT Claims 结构体
type JwtClaims struct {
	UserId      uint   `json:"userId"`
	Username    string `json:"username"`
	AuthorityId uint   `json:"authorityId"`
	jwt.RegisteredClaims
}

// jwtBlacklist JWT 黑名单模型
type jwtBlacklist struct {
	Jwt string `gorm:"type:text;column:jwt"`
}

// TableName 指定表名
func (jwtBlacklist) TableName() string {
	return "jwt_blacklists"
}

// JwtMiddleware JWT 中间件
type JwtMiddleware struct {
	accessSecret string
	accessExpire int64
	db           *gorm.DB
}

// NewJwtMiddleware 创建 JWT 中间件
func NewJwtMiddleware(accessSecret string, accessExpire int64, db *gorm.DB) rest.Middleware {
	jm := &JwtMiddleware{
		accessSecret: accessSecret,
		accessExpire: accessExpire,
		db:           db,
	}
	return jm.Handle
}

// Handle 中间件处理函数
func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 从 Authorization: Bearer <token> 读取 token
		token := GetToken(r)
		if token == "" {
			unauthorizedResponse(w, "未登录或非法访问，请登录")
			return
		}

		// 2. 检查 token 是否在黑名单
		if m.isBlacklist(r.Context(), token) {
			unauthorizedResponse(w, "您的帐户异地登陆或令牌失效")
			return
		}

		// 3. 解析 JWT token
		claims, err := m.parseToken(token)
		if err != nil {
			logx.WithContext(r.Context()).Errorw("JWT 解析失败",
				logx.Field("error", err.Error()),
				logx.Field("module", "jwt"),
				logx.Field("action", "parse_token"),
			)
			unauthorizedResponse(w, "登录已过期，请重新登录")
			return
		}

		// 4. 如果 token 即将过期（在 BufferTime 内），自动续期
		ctx := r.Context()
		if claims.ExpiresAt != nil && claims.ExpiresAt.Unix()-time.Now().Unix() < bufferTime {
			newToken, newExpiresAt, renewErr := m.renewToken(claims)
			if renewErr != nil {
				logx.WithContext(ctx).Errorw("token 续期失败",
					logx.Field("error", renewErr.Error()),
					logx.Field("user_id", claims.UserId),
					logx.Field("module", "jwt"),
					logx.Field("action", "renew_token"),
				)
			} else {
				w.Header().Set("new-token", newToken)
				w.Header().Set("new-expires-at", strconv.FormatInt(newExpiresAt, 10))
			}
		}

		// 5. 将 claims 信息注入到 context 中
		ctx = context.WithValue(ctx, UserIdKey, claims.UserId)
		ctx = context.WithValue(ctx, UsernameKey, claims.Username)
		ctx = context.WithValue(ctx, AuthorityIdKey, claims.AuthorityId)

		// 6. 调用 next(w, r)
		next(w, r.WithContext(ctx))
	}
}

// GetToken 从请求头 Authorization 中提取 Bearer token。
// 供需要原始 token 的接口（如 JWT 拉黑）复用，等价于原 Gin 后端的 utils.GetToken。
func GetToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return ""
	}
	return parts[1]
}

// isBlacklist 检查 token 是否在黑名单中
func (m *JwtMiddleware) isBlacklist(ctx context.Context, token string) bool {
	var count int64
	if err := m.db.WithContext(ctx).Model(&jwtBlacklist{}).
		Where("jwt = ?", token).Count(&count).Error; err != nil {
		logx.WithContext(ctx).Errorw("查询 JWT 黑名单失败",
			logx.Field("error", err.Error()),
			logx.Field("module", "jwt"),
			logx.Field("action", "check_blacklist"),
		)
		return false
	}
	return count > 0
}

// parseToken 解析 JWT token
func (m *JwtMiddleware) parseToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.accessSecret), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrap(err, "token 已过期")
		}
		return nil, errors.Wrap(err, "token 解析失败")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("无效的 token")
	}

	return claims, nil
}

// createToken 创建新的 JWT token
func (m *JwtMiddleware) createToken(claims *JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(m.accessSecret))
	if err != nil {
		return "", errors.Wrap(err, "token 签名失败")
	}
	return signedToken, nil
}

// renewToken 续期 token
func (m *JwtMiddleware) renewToken(oldClaims *JwtClaims) (string, int64, error) {
	newExpiresAt := time.Now().Add(time.Duration(m.accessExpire) * time.Second)
	oldClaims.ExpiresAt = jwt.NewNumericDate(newExpiresAt)
	oldClaims.NotBefore = jwt.NewNumericDate(time.Now().Add(-1000))

	newToken, err := m.createToken(oldClaims)
	if err != nil {
		return "", 0, errors.Wrap(err, "创建新 token 失败")
	}

	return newToken, newExpiresAt.Unix(), nil
}

// unauthorizedResponse 返回 401 响应
func unauthorizedResponse(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"code":401,"data":null,"msg":"` + msg + `"}`))
}
