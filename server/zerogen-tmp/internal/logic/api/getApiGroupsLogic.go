// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiGroupsLogic {
	return &GetApiGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiGroupsLogic) GetApiGroups() (resp *types.ApiGroups, err error) {
	// todo: add your logic here and delete this line

	return
}
