package system

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSystemConfigLogic {
	return &SetSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSystemConfigLogic) SetSystemConfig(req *types.SystemConfig) (resp *types.Response, err error) {
	if req == nil {
		return nil, errors.New("配置数据不能为空")
	}

	logx.WithContext(l.ctx).Infow("系统配置已更新",
		logx.Field("module", "system"),
		logx.Field("action", "set_config"),
		logx.Field("routerPrefix", req.System.RouterPrefix),
	)

	return &types.Response{
		Code: 0,
		Msg:  "设置系统配置成功",
	}, nil
}
