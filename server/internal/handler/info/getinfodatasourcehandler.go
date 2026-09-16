package info

import (
	"net/http"

	"zero/internal/logic/info"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetInfoDataSourceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewGetInfoDataSourceLogic(r.Context(), svcCtx)
		resp, err := l.GetInfoDataSource()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
