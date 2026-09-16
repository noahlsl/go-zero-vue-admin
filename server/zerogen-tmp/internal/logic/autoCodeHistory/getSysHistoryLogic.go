// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCodeHistory

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysHistoryLogic {
	return &GetSysHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysHistoryLogic) GetSysHistory(req *types.GetHistoryListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
