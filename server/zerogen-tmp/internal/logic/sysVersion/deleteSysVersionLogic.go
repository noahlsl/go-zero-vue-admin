// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysVersion

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysVersionLogic {
	return &DeleteSysVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysVersionLogic) DeleteSysVersion(req *types.DeleteSysVersionReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
