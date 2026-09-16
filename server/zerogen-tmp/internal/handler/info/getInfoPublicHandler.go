// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/info"
	"zero/zerogen-tmp/internal/svc"
)

func GetInfoPublicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewGetInfoPublicLogic(r.Context(), svcCtx)
		resp, err := l.GetInfoPublic()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
