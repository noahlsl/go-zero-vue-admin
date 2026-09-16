// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/info"
	"zero/zerogen-tmp/internal/svc"
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
