package base

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// milliSecondsPerSecond 秒到毫秒的换算基数，
// 前端按毫秒解析 expiresAt，与原 Gin 后端 claims.ExpiresAt.Unix()*1000 保持一致
const milliSecondsPerSecond = 1000

// loginClaims JWT 声明结构体
type loginClaims struct {
	UserId      uint   `json:"userId"`
	Username    string `json:"username"`
	AuthorityId uint   `json:"authorityId"`
	jwt.RegisteredClaims
}

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登录：校验验证码 → 查用户 → 校验密码 → 签发 JWT → 写登录日志
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 1. 校验验证码
	if req.CaptchaId != "" && req.Captcha != "" {
		if !store.Verify(req.CaptchaId, req.Captcha, true) {
			l.recordLoginLog(req.Username, "", "", false, "验证码错误")
			return nil, errors.New("验证码错误")
		}
	}

	// 2. 根据用户名查询用户
	user, err := l.findUserByUsername(req.Username)
	if err != nil {
		l.recordLoginLog(req.Username, "", "", false, "用户名不存在或者密码错误")
		return nil, errors.New("用户名不存在或者密码错误")
	}

	// 3. 校验密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		l.recordLoginLog(req.Username, "", "", false, "用户名不存在或者密码错误")
		return nil, errors.New("用户名不存在或者密码错误")
	}

	// 4. 校验角色默认路由，无效时回退 404，避免前端登录后跳转首页失败
	conv.EnsureDefaultRouter(l.ctx, l.svcCtx.DB, &user.Authority)

	// 5. 检查用户是否被冻结
	if user.Enable != 1 {
		l.recordLoginLog(req.Username, "", "", false, "用户被禁止登录")
		return nil, errors.New("用户被禁止登录")
	}

	// 6. 签发 JWT
	token, expiresAt, err := l.createToken(user)
	if err != nil {
		l.recordLoginLog(req.Username, "", "", false, "获取token失败")
		return nil, errors.Wrap(err, "获取token失败")
	}

	// 7. 记录登录成功日志
	l.recordLoginLog(req.Username, "", "", true, "登录成功")

	// 8. 组装响应（包含角色与角色列表，密码不返回）
	resp = &types.LoginResp{
		User:      conv.ToTypesUser(*user),
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return resp, nil
}

// findUserByUsername 根据用户名查询用户，同时预加载当前角色与角色列表
func (l *LoginLogic) findUserByUsername(username string) (*model.SysUser, error) {
	var user model.SysUser
	err := l.svcCtx.DB.WithContext(l.ctx).
		Preload("Authorities").Preload("Authority").
		Where("username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.Wrap(err, "查询用户失败")
	}
	return &user, nil
}

// createToken 生成 JWT Token
func (l *LoginLogic) createToken(user *model.SysUser) (token string, expiresAt int64, err error) {
	expire := time.Now().Add(time.Duration(l.svcCtx.Config.Auth.AccessExpire) * time.Second)
	expiresAt = expire.Unix() * milliSecondsPerSecond

	claims := loginClaims{
		UserId:      user.ID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-zero-vue-admin",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return "", 0, errors.Wrap(err, "签名JWT失败")
	}
	return token, expiresAt, nil
}

// recordLoginLog 异步记录登录日志
func (l *LoginLogic) recordLoginLog(username, ip, agent string, status bool, msg string) {
	loginLog := model.SysLoginLog{
		Username:     username,
		Ip:           ip,
		Agent:        agent,
		Status:       status,
		ErrorMessage: msg,
	}
	if err := l.svcCtx.DB.WithContext(l.ctx).Create(&loginLog).Error; err != nil {
		logx.Errorw("记录登录日志失败",
			logx.Field("error", err.Error()),
			logx.Field("username", username),
			logx.Field("module", "login"),
		)
	}
}
