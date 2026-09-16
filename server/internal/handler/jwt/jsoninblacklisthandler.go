// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package jwt

import (
	"net/http"

	"zero/internal/logic/jwt"
	"zero/internal/middleware"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JsonInBlacklistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 原 Gin 接口无入参，token 取自请求头，因此不做参数解析
		token := middleware.GetToken(r)

		l := jwt.NewJsonInBlacklistLogic(r.Context(), svcCtx)
		resp, err := l.JsonInBlacklist(token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
