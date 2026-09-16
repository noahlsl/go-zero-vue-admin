// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/sysExportTemplate"
	"zero/zerogen-tmp/internal/svc"
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
