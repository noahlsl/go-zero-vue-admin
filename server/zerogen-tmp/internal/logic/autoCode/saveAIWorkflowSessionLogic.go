// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAIWorkflowSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveAIWorkflowSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAIWorkflowSessionLogic {
	return &SaveAIWorkflowSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAIWorkflowSessionLogic) SaveAIWorkflowSession(req *types.SaveAIWorkflowSessionReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
