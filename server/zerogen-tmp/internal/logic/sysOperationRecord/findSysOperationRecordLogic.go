// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysOperationRecord

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysOperationRecordLogic {
	return &FindSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysOperationRecordLogic) FindSysOperationRecord(req *types.GetById) (resp *types.SysOperationRecordRes, err error) {
	// todo: add your logic here and delete this line

	return
}
