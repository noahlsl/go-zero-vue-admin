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

func GetDictionaryTreeListByTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDictionaryTreeListByTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysDictionaryDetail.NewGetDictionaryTreeListByTypeLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryTreeListByType(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
