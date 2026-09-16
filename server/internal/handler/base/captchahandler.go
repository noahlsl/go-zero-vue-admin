// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package base

import (
	"net/http"

	"zero/internal/logic/base"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 原 Gin 接口无入参，前端不传请求体，因此不做参数解析
		l := base.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
