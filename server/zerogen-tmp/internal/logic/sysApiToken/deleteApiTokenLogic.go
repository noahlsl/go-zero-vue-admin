// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysApiToken

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiTokenLogic {
	return &DeleteApiTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiTokenLogic) DeleteApiToken(req *types.DeleteApiTokenReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
