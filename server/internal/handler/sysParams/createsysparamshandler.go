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

func CreateSysParamsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSysParamsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysParams.NewCreateSysParamsLogic(r.Context(), svcCtx)
		resp, err := l.CreateSysParams(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
