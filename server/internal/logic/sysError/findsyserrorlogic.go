package sysError

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysErrorLogic {
	return &FindSysErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysErrorLogic) FindSysError(req *types.GetById) (resp *types.SysErrorRes, err error) {
	// 1. 根据ID查询错误日志
	var sysError model.SysError
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&sysError).Error; err != nil {
		return nil, errors.Wrap(err, "查询错误日志失败")
	}

	return &types.SysErrorRes{
		ErrorInfo: modelSysErrorToTypes(sysError),
	}, nil
}

// modelSysErrorToTypes 将 model.SysError 转换为 types.SysError
func modelSysErrorToTypes(m model.SysError) types.SysError {
	return types.SysError{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		Form:      m.Form,
		Info:      m.Info,
		Level:     m.Level,
		Solution:  m.Solution,
		Status:    m.Status,
	}
}
