package email

import (
	"net/http"

	"zero/internal/logic/email"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
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
