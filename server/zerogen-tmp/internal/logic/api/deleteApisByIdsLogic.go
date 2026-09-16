// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApisByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApisByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApisByIdsLogic {
	return &DeleteApisByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApisByIdsLogic) DeleteApisByIds(req *types.IdsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
