// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authority

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleUsersLogic {
	return &SetRoleUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleUsersLogic) SetRoleUsers(req *types.SetRoleUsersReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
