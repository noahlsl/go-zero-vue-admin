// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCodeHistory

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysHistoryLogic {
	return &DelSysHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysHistoryLogic) DelSysHistory(req *types.DeleteHistoryReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
