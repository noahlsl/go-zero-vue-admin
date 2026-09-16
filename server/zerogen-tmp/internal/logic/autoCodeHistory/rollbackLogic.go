// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCodeHistory

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackLogic {
	return &RollbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RollbackLogic) Rollback(req *types.RollBackReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
