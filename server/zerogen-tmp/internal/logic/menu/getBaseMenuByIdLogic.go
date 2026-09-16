// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package menu

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseMenuByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseMenuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseMenuByIdLogic {
	return &GetBaseMenuByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBaseMenuByIdLogic) GetBaseMenuById(req *types.GetById) (resp *types.SysMenu, err error) {
	// todo: add your logic here and delete this line

	return
}
