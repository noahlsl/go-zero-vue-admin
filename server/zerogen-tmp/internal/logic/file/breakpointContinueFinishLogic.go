// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package file

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BreakpointContinueFinishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBreakpointContinueFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BreakpointContinueFinishLogic {
	return &BreakpointContinueFinishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BreakpointContinueFinishLogic) BreakpointContinueFinish(req *types.BreakpointContinueFinishReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
