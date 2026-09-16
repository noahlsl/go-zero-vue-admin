// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package casbin

import (
	"net/http"

	"zero/internal/logic/casbin"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPolicyPathByAuthorityIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPolicyPathByAuthorityIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := casbin.NewGetPolicyPathByAuthorityIdLogic(r.Context(), svcCtx)
		resp, err := l.GetPolicyPathByAuthorityId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
