package info

import (
	"net/http"

	"zero/internal/logic/info"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
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
