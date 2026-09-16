// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package system

import (
	"net/http"

	"zero/internal/logic/system"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetServerInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := system.NewGetServerInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetServerInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
