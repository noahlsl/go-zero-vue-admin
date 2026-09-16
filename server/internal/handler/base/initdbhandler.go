// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package base

import (
	"net/http"

	"zero/internal/logic/base"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func InitDBHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewInitDBLogic(r.Context(), svcCtx)
		err := l.InitDB()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
