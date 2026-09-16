package sysDictionaryDetail

import (
	"net/http"

	"zero/internal/logic/sysDictionaryDetail"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDictionaryDetailsByParentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDictionaryDetailsByParentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysDictionaryDetail.NewGetDictionaryDetailsByParentLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryDetailsByParent(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
