// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package file

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BreakpointContinueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBreakpointContinueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BreakpointContinueLogic {
	return &BreakpointContinueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BreakpointContinueLogic) BreakpointContinue(req *types.BreakpointContinueReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
