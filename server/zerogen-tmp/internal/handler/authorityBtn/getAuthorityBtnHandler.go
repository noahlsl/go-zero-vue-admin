// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authorityBtn

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/authorityBtn"
	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"
)

func GetAuthorityBtnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAuthorityBtnReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authorityBtn.NewGetAuthorityBtnLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthorityBtn(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
