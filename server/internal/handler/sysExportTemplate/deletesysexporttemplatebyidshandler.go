// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"net/http"

	"zero/internal/logic/sysExportTemplate"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteSysExportTemplateByIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysExportTemplate.NewDeleteSysExportTemplateByIdsLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSysExportTemplateByIds(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
