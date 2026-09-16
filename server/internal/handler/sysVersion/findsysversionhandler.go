// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysVersion

import (
	"net/http"

	"zero/internal/logic/sysVersion"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindSysVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetById
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysVersion.NewFindSysVersionLogic(r.Context(), svcCtx)
		resp, err := l.FindSysVersion(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
