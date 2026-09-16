// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/base"
	"zero/zerogen-tmp/internal/svc"
)

func CheckDBHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewCheckDBLogic(r.Context(), svcCtx)
		err := l.CheckDB()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
