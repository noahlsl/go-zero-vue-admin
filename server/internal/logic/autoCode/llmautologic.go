package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LLMAutoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLLMAutoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LLMAutoLogic {
	return &LLMAutoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LLMAuto 调用大模型自动生成代码
// TODO: 实现完整的大模型调用逻辑，需要配置 LLM 服务地址
func (l *LLMAutoLogic) LLMAuto() (resp *types.Response, err error) {
	return &types.Response{
		Code: 0,
		Data: nil,
		Msg:  fmt.Sprintf("LLM 自动生成功能暂未配置"),
	}, nil
}
