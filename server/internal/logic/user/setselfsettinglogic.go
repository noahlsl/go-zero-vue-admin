package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetSelfSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSelfSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSelfSettingLogic {
	return &SetSelfSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSelfSettingLogic) SetSelfSetting(setting map[string]any) error {
	// 1. 从 context 获取当前登录用户 ID
	userId := middleware.GetUserId(l.ctx)
	if userId == 0 {
		return errors.New("未获取到用户信息")
	}

	// 2. 更新用户个人配置
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).Where("id = ?", userId).
		Update("origin_setting", setting).Error; err != nil {
		l.Logger.Errorw("设置用户个人配置失败",
			logx.Field("error", err),
			logx.Field("user_id", userId),
			logx.Field("module", "user"),
			logx.Field("action", "set_self_setting"),
		)
		return errors.Wrap(err, "设置用户个人配置失败")
	}

	return nil
}
