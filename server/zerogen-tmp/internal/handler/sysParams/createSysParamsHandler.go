// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/sysParams"
	"zero/zerogen-tmp/internal/svc"
)

func CreateSysParamsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysParams.NewCreateSysParamsLogic(r.Context(), svcCtx)
		resp, err := l.CreateSysParams()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
