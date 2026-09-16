package sysError

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysErrorLogic {
	return &DeleteSysErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysErrorLogic) DeleteSysError(req *types.DeleteSysErrorReq) (resp *types.Response, err error) {
	// 1. 校验记录是否存在
	var sysError model.SysError
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&sysError).Error; err != nil {
		return nil, errors.Wrap(err, "错误日志记录不存在")
	}

	// 2. 执行软删除
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).Delete(&model.SysError{}).Error; err != nil {
		l.Logger.Errorw("删除错误日志失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysError"),
			logx.Field("action", "delete"),
		)
		return nil, errors.Wrap(err, "删除错误日志失败")
	}

	return &types.Response{Msg: "删除成功"}, nil
}
