// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"net/http"

	"zero/internal/logic/sysExportTemplate"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSysExportTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysExportTemplate.NewUpdateSysExportTemplateLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSysExportTemplate()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
