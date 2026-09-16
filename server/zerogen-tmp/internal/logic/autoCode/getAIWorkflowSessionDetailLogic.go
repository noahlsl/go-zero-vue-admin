// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAIWorkflowSessionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAIWorkflowSessionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAIWorkflowSessionDetailLogic {
	return &GetAIWorkflowSessionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAIWorkflowSessionDetailLogic) GetAIWorkflowSessionDetail(req *types.GetById) (resp *types.AIWorkflowSessionDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
