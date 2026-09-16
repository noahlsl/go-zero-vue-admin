package api

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiLogic) DeleteApi(req *types.GetById) (resp *types.Response, err error) {
	var apiModel model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).First(&apiModel, "id = ?", req.ID).Error; err != nil {
		return nil, errors.Wrap(err, "查询API失败")
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Delete(&apiModel).Error; err != nil {
		return nil, errors.Wrap(err, "删除API失败")
	}

	// 清除 Casbin 中对应 path + method 的策略
	if err = l.clearCasbinByPath(apiModel.Path, apiModel.Method); err != nil {
		return nil, errors.Wrap(err, "清除Casbin策略失败")
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}

func (l *DeleteApiLogic) clearCasbinByPath(path, method string) error {
	return l.svcCtx.DB.WithContext(l.ctx).
		Where("v1 = ? AND v2 = ?", path, method).
		Delete(&model.SysCasbinRule{}).
		Error
}
