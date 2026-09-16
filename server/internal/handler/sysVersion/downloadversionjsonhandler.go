// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysVersion

import (
	"net/http"

	"zero/internal/logic/sysVersion"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadVersionJsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysVersion.NewDownloadVersionJsonLogic(r.Context(), svcCtx)
		resp, err := l.DownloadVersionJson()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
