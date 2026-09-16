// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysVersion

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysVersionLogic {
	return &FindSysVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysVersionLogic) FindSysVersion(req *types.GetById) (resp *types.SysVersionRes, err error) {
	// todo: add your logic here and delete this line

	return
}
