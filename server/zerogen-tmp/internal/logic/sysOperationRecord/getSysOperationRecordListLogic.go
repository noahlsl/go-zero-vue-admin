// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysOperationRecord

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysOperationRecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysOperationRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysOperationRecordListLogic {
	return &GetSysOperationRecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysOperationRecordListLogic) GetSysOperationRecordList(req *types.GetSysOperationRecordListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
