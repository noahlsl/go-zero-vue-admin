// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiRolesLogic {
	return &GetApiRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiRolesLogic) GetApiRoles(req *types.GetApiRolesReq) (resp []uint, err error) {
	// todo: add your logic here and delete this line

	return
}
