// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysOperationRecord

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysOperationRecordLogic {
	return &DeleteSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysOperationRecordLogic) DeleteSysOperationRecord(req *types.DeleteSysOperationRecordReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
