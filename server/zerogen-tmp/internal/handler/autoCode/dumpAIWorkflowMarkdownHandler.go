// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/autoCode"
	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"
)

func DumpAIWorkflowMarkdownHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DumpAIWorkflowMarkdownReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := autoCode.NewDumpAIWorkflowMarkdownLogic(r.Context(), svcCtx)
		resp, err := l.DumpAIWorkflowMarkdown(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
