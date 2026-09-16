// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"net/http"

	"zero/internal/logic/sysParams"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSysParamsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSysParamsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysParams.NewUpdateSysParamsLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSysParams(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
