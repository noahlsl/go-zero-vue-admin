package file

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

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
	// TODO: 实现断点续传切片上传逻辑
	return &types.Response{Msg: "切片上传成功"}, nil
}
