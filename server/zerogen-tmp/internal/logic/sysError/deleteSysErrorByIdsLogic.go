// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysError

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysErrorByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysErrorByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysErrorByIdsLogic {
	return &DeleteSysErrorByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysErrorByIdsLogic) DeleteSysErrorByIds(req *types.DeleteSysErrorByIdsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
