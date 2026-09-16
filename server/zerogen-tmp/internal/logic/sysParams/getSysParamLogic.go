// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysParamLogic {
	return &GetSysParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysParamLogic) GetSysParam(req *types.GetById) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
