// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

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

func (l *LLMAutoSSELogic) LLMAutoSSE() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
