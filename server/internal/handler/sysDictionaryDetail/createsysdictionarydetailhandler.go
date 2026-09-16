// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"net/http"

	"zero/internal/logic/sysDictionaryDetail"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateSysDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSysDictionaryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysDictionaryDetail.NewCreateSysDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.CreateSysDictionaryDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
