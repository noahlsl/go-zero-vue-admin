// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type McpRoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMcpRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *McpRoutesLogic {
	return &McpRoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *McpRoutesLogic) McpRoutes() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
