// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authorityBtn

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorityBtnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthorityBtnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorityBtnLogic {
	return &GetAuthorityBtnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthorityBtnLogic) GetAuthorityBtn(req *types.GetAuthorityBtnReq) (resp *types.AuthorityBtnRes, err error) {
	// todo: add your logic here and delete this line

	return
}
