// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCodeHistory

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMetaLogic {
	return &GetMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMetaLogic) GetMeta(req *types.CreateHistoryReq) (resp *types.SysAutoCodeHistory, err error) {
	// todo: add your logic here and delete this line

	return
}
