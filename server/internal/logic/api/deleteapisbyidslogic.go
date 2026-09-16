package api

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApisByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApisByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApisByIdsLogic {
	return &DeleteApisByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApisByIdsLogic) DeleteApisByIds(req *types.IdsReq) (resp *types.Response, err error) {
	// 先查询要删除的 API，用于清除 Casbin 策略
	var apis []model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id IN ?", req.Ids).Find(&apis).Error; err != nil {
		return nil, errors.Wrap(err, "查询待删除API失败")
	}

	// 批量删除
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id IN ?", req.Ids).Delete(&model.SysApi{}).Error; err != nil {
		return nil, errors.Wrap(err, "批量删除API失败")
	}

	// 清除 Casbin 策略
	for _, apiItem := range apis {
		if err = l.svcCtx.DB.WithContext(l.ctx).
			Where("v1 = ? AND v2 = ?", apiItem.Path, apiItem.Method).
			Delete(&model.SysCasbinRule{}).Error; err != nil {
			return nil, errors.Wrap(err, "清除Casbin策略失败")
		}
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}
