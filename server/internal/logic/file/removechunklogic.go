package file

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveChunkLogic {
	return &RemoveChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveChunkLogic) RemoveChunk(req *types.RemoveChunkReq) (resp *types.Response, err error) {
	// TODO: 实现删除切片逻辑
	return &types.Response{Msg: "切片删除成功"}, nil
}
