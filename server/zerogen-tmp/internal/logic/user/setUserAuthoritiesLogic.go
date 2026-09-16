// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserAuthoritiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserAuthoritiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserAuthoritiesLogic {
	return &SetUserAuthoritiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserAuthoritiesLogic) SetUserAuthorities(req *types.SetUserAuthoritiesReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
