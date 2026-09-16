package middleware

import (
	"net/http"

	"zero/internal/config"
)

const (
	defaultAllowHeaders  = "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id"
	defaultAllowMethods  = "POST, GET, OPTIONS, DELETE, PUT"
	defaultExposeHeaders = "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type,New-Token,New-Expires-At"
)

// CorsMiddleware CORS 中间件
type CorsMiddleware struct {
	whitelist []config.CORSWhitelist
	mode      string // "allow-all" 或 "strict-whitelist"
}

// NewCorsMiddleware 创建 CORS 中间件
func NewCorsMiddleware(cfg config.CORS) *CorsMiddleware {
	return &CorsMiddleware{
		whitelist: cfg.Whitelist,
		mode:      cfg.Mode,
	}
}

// Handle 处理跨域请求
func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if m.mode == "allow-all" {
			setDefaultCORSHeaders(w, origin)
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next(w, r)
			return
		}

		// 白名单模式
		wl := matchWhitelist(origin, m.whitelist)
		if wl != nil {
			setWhitelistCORSHeaders(w, wl)
		}

		// 严格白名单模式：未匹配且非健康检查则拒绝
		if wl == nil && m.mode == "strict-whitelist" && !(r.Method == http.MethodGet && r.URL.Path == "/health") {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

// setDefaultCORSHeaders 设置全放行模式的 CORS 响应头
func setDefaultCORSHeaders(w http.ResponseWriter, origin string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", defaultAllowHeaders)
	w.Header().Set("Access-Control-Allow-Methods", defaultAllowMethods)
	w.Header().Set("Access-Control-Expose-Headers", defaultExposeHeaders)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

// setWhitelistCORSHeaders 根据白名单设置 CORS 响应头
func setWhitelistCORSHeaders(w http.ResponseWriter, wl *config.CORSWhitelist) {
	w.Header().Set("Access-Control-Allow-Origin", wl.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", wl.AllowHeaders)
	w.Header().Set("Access-Control-Allow-Methods", wl.AllowMethods)
	w.Header().Set("Access-Control-Expose-Headers", wl.ExposeHeaders)
	if wl.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

// matchWhitelist 在白名单中匹配当前请求的 Origin
func matchWhitelist(currentOrigin string, whitelist []config.CORSWhitelist) *config.CORSWhitelist {
	for i := range whitelist {
		if currentOrigin == whitelist[i].AllowOrigin {
			return &whitelist[i]
		}
	}
	return nil
}
