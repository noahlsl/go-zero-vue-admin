// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysOperationRecord

import (
	"net/http"

	"zero/internal/logic/sysOperationRecord"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysOperationRecordListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSysOperationRecordListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysOperationRecord.NewGetSysOperationRecordListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysOperationRecordList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
