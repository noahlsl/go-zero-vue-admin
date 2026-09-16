// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package email

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/email"
	"zero/zerogen-tmp/internal/svc"
)

func EmailTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := email.NewEmailTestLogic(r.Context(), svcCtx)
		resp, err := l.EmailTest()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
