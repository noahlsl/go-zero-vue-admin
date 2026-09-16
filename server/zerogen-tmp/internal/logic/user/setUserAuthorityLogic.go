// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserAuthorityLogic {
	return &SetUserAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserAuthorityLogic) SetUserAuthority(req *types.SetUserAuthReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
