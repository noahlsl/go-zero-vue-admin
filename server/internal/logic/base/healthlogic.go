package base

import (
	"context"

	"zero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Health 检查服务健康状态，验证数据库连接是否正常
func (l *HealthLogic) Health() error {
	sqlDB, err := l.svcCtx.DB.DB()
	if err != nil {
		logx.Errorw("健康检查失败：获取数据库连接出错", logx.Field("error", err.Error()))
		return err
	}

	if err = sqlDB.PingContext(l.ctx); err != nil {
		logx.Errorw("健康检查失败：数据库 Ping 不通", logx.Field("error", err.Error()))
		return err
	}

	return nil
}
