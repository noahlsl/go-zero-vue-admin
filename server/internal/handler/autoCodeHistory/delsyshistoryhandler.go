// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCodeHistory

import (
	"net/http"

	"zero/internal/logic/autoCodeHistory"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelSysHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := autoCodeHistory.NewDelSysHistoryLogic(r.Context(), svcCtx)
		resp, err := l.DelSysHistory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
