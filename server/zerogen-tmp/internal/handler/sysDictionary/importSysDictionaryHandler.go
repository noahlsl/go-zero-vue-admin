// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/zerogen-tmp/internal/logic/sysDictionary"
	"zero/zerogen-tmp/internal/svc"
)

func ImportSysDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysDictionary.NewImportSysDictionaryLogic(r.Context(), svcCtx)
		err := l.ImportSysDictionary()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
