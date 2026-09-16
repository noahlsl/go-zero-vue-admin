package api

import (
	"context"

	"zero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IgnoreApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIgnoreApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IgnoreApiLogic {
	return &IgnoreApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IgnoreApiLogic) IgnoreApi() error {
	// 注意：原 Gin 项目接收 SysIgnoreApi body 参数（含 path, method, flag）
	// go-zero handler 未解析请求体，此接口暂为空操作
	// 如需忽略 API 功能，需修改 handler 添加请求体解析

	return nil
}
