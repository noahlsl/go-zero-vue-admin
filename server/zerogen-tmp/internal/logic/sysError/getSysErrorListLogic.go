// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysError

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysErrorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysErrorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysErrorListLogic {
	return &GetSysErrorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysErrorListLogic) GetSysErrorList(req *types.GetSysErrorListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
