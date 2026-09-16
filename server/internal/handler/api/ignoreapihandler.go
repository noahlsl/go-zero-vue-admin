// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"net/http"

	"zero/internal/logic/api"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func IgnoreApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := api.NewIgnoreApiLogic(r.Context(), svcCtx)
		err := l.IgnoreApi()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
