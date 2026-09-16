package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// CasbinHandlerMiddleware Casbin 权限校验中间件
// 职责：
// 1. 从 context 中读取 JWT 中间件注入的用户信息（userId、authorityId）
// 2. 查询角色的 API 权限列表
// 3. 匹配当前请求的 method + path
// 4. 有权限则放行，否则返回 403
type CasbinHandlerMiddleware struct {
	// TODO: 后续集成 casbin 后添加 enforcer 字段
	// enforcer *casbin.SyncedCachedEnforcer
}

// NewCasbinHandlerMiddleware 创建 CasbinHandler 中间件实例
func NewCasbinHandlerMiddleware() *CasbinHandlerMiddleware {
	return &CasbinHandlerMiddleware{}
}

// contextKey 自定义 context key 类型，避免冲突
type contextKey string

const (
	// UserIdKey 用户ID在context中的key
	UserIdKey contextKey = "userId"
	// AuthorityIdKey 角色ID在context中的key
	AuthorityIdKey contextKey = "authorityId"
	// UsernameKey 用户名在context中的key
	UsernameKey contextKey = "username"
)

// Handle 中间件处理函数
func (m *CasbinHandlerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 从 context 中获取 JWT 中间件注入的用户信息
		userId, _ := r.Context().Value(UserIdKey).(uint)
		authorityId, _ := r.Context().Value(AuthorityIdKey).(uint)

		// 兜底校验：如果用户信息为空，说明 JWT 中间件未通过或 token 无效
		if userId == 0 {
			http.Error(w, "未授权访问", http.StatusUnauthorized)
			return
		}

		// 2. 获取请求路径和方法
		path := r.URL.Path
		method := r.Method

		// 3. Casbin 策略校验
		// TODO: 集成 casbin 后启用完整的策略校验
		// 目前先实现基础的用户信息提取，策略校验待 casbin 集成后完善
		/*
			authorityIdStr := strconv.Itoa(int(authorityId))
			success, err := m.enforcer.Enforce(authorityIdStr, path, method)
			if err != nil {
				logx.WithContext(r.Context()).Errorw("Casbin 策略校验异常",
					logx.Field("error", err.Error()),
					logx.Field("authority_id", authorityIdStr),
					logx.Field("path", path),
					logx.Field("method", method),
					logx.Field("module", "casbin"),
					logx.Field("action", "enforce"),
				)
				http.Error(w, "权限校验异常", http.StatusInternalServerError)
				return
			}
			if !success {
				logx.WithContext(r.Context()).Infow("权限不足",
					logx.Field("authority_id", authorityIdStr),
					logx.Field("path", path),
					logx.Field("method", method),
					logx.Field("module", "casbin"),
					logx.Field("action", "deny"),
				)
				http.Error(w, "权限不足", http.StatusForbidden)
				return
			}
		*/

		// 4. 记录请求信息，便于调试
		logx.WithContext(r.Context()).Debugw("CasbinHandler 请求信息",
			logx.Field("user_id", userId),
			logx.Field("authority_id", authorityId),
			logx.Field("path", path),
			logx.Field("method", method),
			logx.Field("module", "casbin"),
			logx.Field("action", "request_info"),
		)

		// 5. 放行到下一个 handler（context 已由 JWT 中间件注入用户信息）
		next(w, r)
	}
}

// GetUserId 从 context 中获取用户ID
func GetUserId(ctx context.Context) uint {
	if userId, ok := ctx.Value(UserIdKey).(uint); ok {
		return userId
	}
	return 0
}

// GetAuthorityId 从 context 中获取角色ID
func GetAuthorityId(ctx context.Context) uint {
	if authorityId, ok := ctx.Value(AuthorityIdKey).(uint); ok {
		return authorityId
	}
	return 0
}

// GetUsername 从 context 中获取用户名
func GetUsername(ctx context.Context) string {
	if username, ok := ctx.Value(UsernameKey).(string); ok {
		return username
	}
	return ""
}

// formatAuthorityId 格式化 authorityId 为字符串
// 用于 casbin 策略匹配
func formatAuthorityId(authorityId uint) string {
	return strconv.Itoa(int(authorityId))
}
