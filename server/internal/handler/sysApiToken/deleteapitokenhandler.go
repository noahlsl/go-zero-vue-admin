// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysApiToken

import (
	"net/http"

	"zero/internal/logic/sysApiToken"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteApiTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteApiTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysApiToken.NewDeleteApiTokenLogic(r.Context(), svcCtx)
		resp, err := l.DeleteApiToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
