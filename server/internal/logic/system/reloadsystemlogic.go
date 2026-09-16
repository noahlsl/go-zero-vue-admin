package system

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReloadSystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReloadSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReloadSystemLogic {
	return &ReloadSystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReloadSystemLogic) ReloadSystem() (resp *types.Response, err error) {
	logx.WithContext(l.ctx).Infow("系统重载成功",
		logx.Field("module", "system"),
		logx.Field("action", "reload"),
	)

	return &types.Response{
		Code: 0,
		Msg:  "重新加载系统成功",
	}, nil
}
