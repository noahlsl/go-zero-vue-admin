// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package file

import (
	"net/http"

	"zero/internal/logic/file"
	"zero/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewUploadFileLogic(file.WithHTTPRequest(r.Context(), r), svcCtx)
		resp, err := l.UploadFile()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
