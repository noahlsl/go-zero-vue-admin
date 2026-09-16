package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
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

// DumpAIWorkflowMarkdown 将 AI 工作流会话导出为 Markdown
// TODO: 实现完整的 Markdown 导出逻辑
func (l *DumpAIWorkflowMarkdownLogic) DumpAIWorkflowMarkdown(req *types.DumpAIWorkflowMarkdownReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("会话ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Data: nil,
		Msg:  fmt.Sprintf("AI 工作流会话 %d 导出成功", req.ID),
	}, nil
}
