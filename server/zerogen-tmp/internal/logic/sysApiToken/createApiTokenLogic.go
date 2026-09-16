// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysApiToken

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiTokenLogic {
	return &CreateApiTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiTokenLogic) CreateApiToken(req *types.CreateApiTokenReq) (resp *types.CreateApiTokenRes, err error) {
	// todo: add your logic here and delete this line

	return
}
