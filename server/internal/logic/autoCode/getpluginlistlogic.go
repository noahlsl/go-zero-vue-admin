package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPluginListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPluginListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPluginListLogic {
	return &GetPluginListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetPluginList 获取插件列表
// TODO: 实现完整的插件列表查询逻辑
func (l *GetPluginListLogic) GetPluginList() (resp *types.Response, err error) {
	return &types.Response{
		Code: 0,
		Data: []interface{}{},
		Msg:  fmt.Sprintf("获取插件列表成功"),
	}, nil
}
