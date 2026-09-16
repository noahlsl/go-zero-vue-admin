// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DumpAIWorkflowMarkdownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDumpAIWorkflowMarkdownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DumpAIWorkflowMarkdownLogic {
	return &DumpAIWorkflowMarkdownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DumpAIWorkflowMarkdownLogic) DumpAIWorkflowMarkdown(req *types.DumpAIWorkflowMarkdownReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
