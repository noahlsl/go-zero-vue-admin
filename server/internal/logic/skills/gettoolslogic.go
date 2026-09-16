package skills

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetToolsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 支持的工具列表
var skillToolOrder = []string{"copilot", "claude", "cursor", "trae", "codex"}
var skillToolLabels = map[string]string{
	"copilot": "Copilot",
	"claude":  "Claude",
	"trae":    "Trae",
	"codex":   "Codex",
	"cursor":  "Cursor",
}

func NewGetToolsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetToolsLogic {
	return &GetToolsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetToolsLogic) GetTools() (resp *types.Response, err error) {
	tools := make([]map[string]string, 0, len(skillToolOrder))
	for _, key := range skillToolOrder {
		tools = append(tools, map[string]string{
			"key":   key,
			"label": skillToolLabels[key],
		})
	}

	return &types.Response{
		Code: 0,
		Data: tools,
		Msg:  fmt.Sprintf("获取工具列表成功，共 %d 个工具", len(tools)),
	}, nil
}
