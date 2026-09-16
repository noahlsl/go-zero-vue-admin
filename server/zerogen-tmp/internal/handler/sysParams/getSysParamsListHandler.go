// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/sysParams"
	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"
)

func GetSysParamsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysParams.NewGetSysParamsListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysParamsList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
