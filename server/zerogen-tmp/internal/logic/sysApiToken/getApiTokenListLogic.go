// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysApiToken

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiTokenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiTokenListLogic {
	return &GetApiTokenListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiTokenListLogic) GetApiTokenList(req *types.GetApiTokenListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
