// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysError

import (
	"net/http"

	"zero/internal/logic/sysError"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysErrorListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSysErrorListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysError.NewGetSysErrorListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysErrorList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
