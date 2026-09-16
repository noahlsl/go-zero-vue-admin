// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysVersion

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysVersionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysVersionListLogic {
	return &GetSysVersionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysVersionListLogic) GetSysVersionList(req *types.GetSysVersionListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
