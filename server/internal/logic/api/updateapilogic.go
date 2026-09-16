package api

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApiLogic) UpdateApi(req *types.UpdateApiReq) (resp *types.Response, err error) {
	// 查询原 API
	var oldApi model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).First(&oldApi, "id = ?", req.ID).Error; err != nil {
		return nil, errors.Wrap(err, "查询原API失败")
	}

	// 如果 path 或 method 变更，检查是否与已有的 API 重复
	if oldApi.Path != req.Path || oldApi.Method != req.Method {
		var count int64
		if err = l.svcCtx.DB.WithContext(l.ctx).
			Model(&model.SysApi{}).
			Where("path = ? AND method = ?", req.Path, req.Method).
			Count(&count).Error; err != nil {
			return nil, errors.Wrap(err, "查询API是否存在失败")
		}
		if count > 0 {
			return nil, errors.New("存在相同api路径")
		}
	}

	// 更新 Casbin 策略中的路径和方法
	if oldApi.Path != req.Path || oldApi.Method != req.Method {
		if err = l.svcCtx.DB.WithContext(l.ctx).
			Model(&model.SysCasbinRule{}).
			Where("v1 = ? AND v2 = ?", oldApi.Path, oldApi.Method).
			Updates(map[string]interface{}{"v1": req.Path, "v2": req.Method}).Error; err != nil {
			return nil, errors.Wrap(err, "更新Casbin策略失败")
		}
	}

	// 更新 API
	apiModel := model.SysApi{
		GvaModel:    model.GvaModel{ID: req.ID},
		Path:        req.Path,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
		Method:      req.Method,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Save(&apiModel).Error; err != nil {
		return nil, errors.Wrap(err, "更新API失败")
	}

	return &types.Response{Code: 0, Msg: "修改成功"}, nil
}
