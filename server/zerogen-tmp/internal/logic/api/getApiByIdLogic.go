// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiByIdLogic) GetApiById(req *types.GetById) (resp *types.SysApi, err error) {
	// todo: add your logic here and delete this line

	return
}
