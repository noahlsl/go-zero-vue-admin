// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package menu

import (
	"net/http"

	"zero/internal/logic/menu"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBaseMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetBaseMenuTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetBaseMenuTree()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
