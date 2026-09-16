package info

import (
	"net/http"

	"zero/internal/logic/info"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindAnnouncementInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := info.NewFindInfoLogic(r.Context(), svcCtx)
		resp, err := l.FindInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
