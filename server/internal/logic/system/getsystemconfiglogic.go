package system

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemConfigLogic {
	return &GetSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemConfigLogic) GetSystemConfig() (resp *types.SystemConfig, err error) {
	resp = &types.SystemConfig{
		System: types.System{
			RouterPrefix:  l.svcCtx.Config.Name,
			UseMultipoint: false,
			ConfigPath:    "",
		},
	}
	return resp, nil
}
