// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"net/http"

	"zero/internal/logic/skills"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGlobalConstraintHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := skills.NewGetGlobalConstraintLogic(r.Context(), svcCtx)
		resp, err := l.GetGlobalConstraint()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
