// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package menu

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMenuRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMenuRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenuRolesLogic {
	return &SetMenuRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMenuRolesLogic) SetMenuRoles(req *types.SetMenuRolesReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
