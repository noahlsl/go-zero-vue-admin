// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package jwt

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/jwt"
	"zero/zerogen-tmp/internal/svc"
)

func JsonInBlacklistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := jwt.NewJsonInBlacklistLogic(r.Context(), svcCtx)
		resp, err := l.JsonInBlacklist()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
