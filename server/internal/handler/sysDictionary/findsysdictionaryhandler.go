// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"net/http"

	"zero/internal/logic/sysDictionary"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindSysDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindSysDictionaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysDictionary.NewFindSysDictionaryLogic(r.Context(), svcCtx)
		resp, err := l.FindSysDictionary(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
