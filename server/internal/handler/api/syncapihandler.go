package api

import (
	"net/http"

	"zero/internal/logic/api"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := api.NewSyncApiLogic(r.Context(), svcCtx)
		resp, err := l.SyncApi()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
