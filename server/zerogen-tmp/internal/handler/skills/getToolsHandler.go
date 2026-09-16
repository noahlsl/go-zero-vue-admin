// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/skills"
	"zero/zerogen-tmp/internal/svc"
)

func GetToolsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := skills.NewGetToolsLogic(r.Context(), svcCtx)
		resp, err := l.GetTools()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
