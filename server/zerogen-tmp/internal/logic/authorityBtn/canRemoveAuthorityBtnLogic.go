// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authorityBtn

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CanRemoveAuthorityBtnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCanRemoveAuthorityBtnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CanRemoveAuthorityBtnLogic {
	return &CanRemoveAuthorityBtnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CanRemoveAuthorityBtnLogic) CanRemoveAuthorityBtn(req *types.GetById) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
