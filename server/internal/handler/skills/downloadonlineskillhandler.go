// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"net/http"

	"zero/internal/logic/skills"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadOnlineSkillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadOnlineSkillReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := skills.NewDownloadOnlineSkillLogic(r.Context(), svcCtx)
		resp, err := l.DownloadOnlineSkill(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
