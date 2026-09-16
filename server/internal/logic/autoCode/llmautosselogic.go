package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LLMAutoSSELogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLLMAutoSSELogic(ctx context.Context, svcCtx *svc.ServiceContext) *LLMAutoSSELogic {
	return &LLMAutoSSELogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LLMAutoSSE 通过 SSE 调用大模型自动生成代码
// TODO: 实现完整的 SSE 流式调用逻辑
func (l *LLMAutoSSELogic) LLMAutoSSE() (resp *types.Response, err error) {
	return &types.Response{
		Code: 0,
		Data: nil,
		Msg:  fmt.Sprintf("LLM SSE 功能暂未配置"),
	}, nil
}
