// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysParamsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysParamsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysParamsListLogic {
	return &GetSysParamsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysParamsListLogic) GetSysParamsList(req *types.PageInfo) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
