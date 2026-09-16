// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"net/http"

	"zero/internal/logic/autoCode"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func McpStopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := autoCode.NewMcpStopLogic(r.Context(), svcCtx)
		resp, err := l.McpStop()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
