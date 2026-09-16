package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type McpStopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMcpStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *McpStopLogic {
	return &McpStopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// McpStop 停止 MCP 服务
// TODO: 实现完整的 MCP 服务停止逻辑
func (l *McpStopLogic) McpStop() (resp *types.Response, err error) {
	return &types.Response{
		Code: 0,
		Msg:  fmt.Sprintf("MCP 服务停止成功"),
	}, nil
}
