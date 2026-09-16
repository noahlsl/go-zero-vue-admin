package base

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckDBLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckDBLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckDBLogic {
	return &CheckDBLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CheckDB 检查数据库是否需要初始化（通过查询用户表判断）
func (l *CheckDBLogic) CheckDB() error {
	var count int64
	err := l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysUser{}).
		Count(&count).Error
	if err != nil {
		logx.Errorw("检查数据库状态失败", logx.Field("error", err.Error()))
		return errors.Wrap(err, "检查数据库状态失败")
	}

	if count == 0 {
		return errors.New("数据库未初始化")
	}

	return nil
}
