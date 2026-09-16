// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package menu

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuAuthorityLogic {
	return &AddMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuAuthorityLogic) AddMenuAuthority(req *types.AddMenuAuthorityReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
