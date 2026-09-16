// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysLoginLog

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLoginLogLogic {
	return &FindLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLoginLogLogic) FindLoginLog(req *types.GetById) (resp *types.SysLoginLogRes, err error) {
	// todo: add your logic here and delete this line

	return
}
