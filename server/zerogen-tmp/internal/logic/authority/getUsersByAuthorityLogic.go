// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authority

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersByAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByAuthorityLogic {
	return &GetUsersByAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersByAuthorityLogic) GetUsersByAuthority(req *types.SetRoleUsersReq) (resp []uint, err error) {
	// todo: add your logic here and delete this line

	return
}
