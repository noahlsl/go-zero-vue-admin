// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysLoginLog

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLoginLogLogic {
	return &DeleteLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLoginLogLogic) DeleteLoginLog(req *types.DeleteLoginLogReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
