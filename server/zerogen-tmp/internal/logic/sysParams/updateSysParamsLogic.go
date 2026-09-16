// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysParams

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysParamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysParamsLogic {
	return &UpdateSysParamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysParamsLogic) UpdateSysParams() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
