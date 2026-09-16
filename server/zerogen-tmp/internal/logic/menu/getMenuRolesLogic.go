// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package menu

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRolesLogic {
	return &GetMenuRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuRolesLogic) GetMenuRoles(req *types.SetMenuRolesReq) (resp *types.MenuRolesRes, err error) {
	// todo: add your logic here and delete this line

	return
}
