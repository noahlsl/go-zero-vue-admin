package sysParams

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysParamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysParamsLogic {
	return &DeleteSysParamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysParamsLogic) DeleteSysParams(req *types.GetById) (resp *types.Response, err error) {
	// 1. 校验记录是否存在
	var sysParams model.SysParams
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&sysParams).Error; err != nil {
		return nil, errors.Wrap(err, "参数记录不存在")
	}

	// 2. 执行软删除
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).Delete(&model.SysParams{}).Error; err != nil {
		l.Logger.Errorw("删除参数失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysParams"),
			logx.Field("action", "delete"),
		)
		return nil, errors.Wrap(err, "删除参数失败")
	}

	return &types.Response{Msg: "删除成功"}, nil
}
