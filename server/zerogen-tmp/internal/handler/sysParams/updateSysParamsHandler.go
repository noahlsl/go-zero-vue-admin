// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/sysParams"
	"zero/zerogen-tmp/internal/svc"
)

func UpdateSysParamsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysParams.NewUpdateSysParamsLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSysParams()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
