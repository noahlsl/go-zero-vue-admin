// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package attachment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/attachment"
	"zero/zerogen-tmp/internal/svc"
)

func GetCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := attachment.NewGetCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
