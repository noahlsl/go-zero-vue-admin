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

func DeleteSysVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSysVersionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysVersion.NewDeleteSysVersionLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSysVersion(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
