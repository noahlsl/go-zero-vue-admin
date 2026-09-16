package sysLoginLog

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLoginLogLogic {
	return &FindLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLoginLogLogic) FindLoginLog(req *types.GetById) (resp *types.SysLoginLogRes, err error) {
	// 1. 根据ID查询登录日志
	var loginLog model.SysLoginLog
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&loginLog).Error; err != nil {
		return nil, errors.Wrap(err, "查询登录日志失败")
	}

	return &types.SysLoginLogRes{
		LoginLog: modelLoginLogToTypes(loginLog),
	}, nil
}

// modelLoginLogToTypes 将 model.SysLoginLog 转换为 types.SysLoginLog
func modelLoginLogToTypes(m model.SysLoginLog) types.SysLoginLog {
	return types.SysLoginLog{
		ID:           m.ID,
		CreatedAt:    m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    m.UpdatedAt.Format(time.RFC3339),
		Username:     m.Username,
		Ip:           m.Ip,
		Agent:        m.Agent,
		Status:       m.Status,
		ErrorMessage: m.ErrorMessage,
		UserId:       m.UserID,
	}
}
