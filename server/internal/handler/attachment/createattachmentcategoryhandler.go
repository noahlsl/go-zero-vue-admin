// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package attachment

import (
	"net/http"

	"zero/internal/logic/attachment"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateAttachmentCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExaAttachmentCategory
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attachment.NewCreateAttachmentCategoryLogic(r.Context(), svcCtx)
		resp, err := l.CreateAttachmentCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
