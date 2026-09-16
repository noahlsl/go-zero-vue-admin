package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type McpListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMcpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *McpListLogic {
	return &McpListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// McpList 获取 MCP 工具列表
// TODO: 实现完整的 MCP 工具列表查询逻辑
func (l *McpListLogic) McpList() (resp *types.Response, err error) {
	return &types.Response{
		Code: 0,
		Data: []interface{}{},
		Msg:  fmt.Sprintf("获取 MCP 工具列表成功"),
	}, nil
}
