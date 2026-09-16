package api

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreshCasbinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreshCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreshCasbinLogic {
	return &FreshCasbinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreshCasbinLogic) FreshCasbin() (resp *types.Response, err error) {
	// go-zero 中 Casbin 策略从数据库实时读取，无需额外刷新缓存操作
	// 原 Gin 项目中的 FreshCasbin 主要刷新内存缓存，此处直接返回成功

	return &types.Response{Code: 0, Msg: "刷新成功"}, nil
}
