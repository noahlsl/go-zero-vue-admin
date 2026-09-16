package sysError

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysErrorLogic {
	return &UpdateSysErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysErrorLogic) UpdateSysError(req *types.UpdateSysErrorReq) (resp *types.Response, err error) {
	// 1. 校验记录是否存在
	var sysError model.SysError
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&sysError).Error; err != nil {
		return nil, errors.Wrap(err, "错误日志记录不存在")
	}

	// 2. 更新记录
	updateMap := map[string]interface{}{
		"form":     req.Form,
		"info":     req.Info,
		"level":    req.Level,
		"solution": req.Solution,
		"status":   req.Status,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysError{}).Where("id = ?", req.ID).Updates(updateMap).Error; err != nil {
		l.Logger.Errorw("更新错误日志失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysError"),
			logx.Field("action", "update"),
		)
		return nil, errors.Wrap(err, "更新错误日志失败")
	}

	return &types.Response{Msg: "更新成功"}, nil
}
