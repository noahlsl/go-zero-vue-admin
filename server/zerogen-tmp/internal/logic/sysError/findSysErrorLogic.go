// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysError

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysErrorLogic {
	return &FindSysErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysErrorLogic) FindSysError(req *types.GetById) (resp *types.SysErrorRes, err error) {
	// todo: add your logic here and delete this line

	return
}
