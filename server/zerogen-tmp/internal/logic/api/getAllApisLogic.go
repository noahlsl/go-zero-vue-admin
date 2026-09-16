// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllApisLogic {
	return &GetAllApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllApisLogic) GetAllApis(req *types.GetAllApisReq) (resp []types.SysApi, err error) {
	// todo: add your logic here and delete this line

	return
}
