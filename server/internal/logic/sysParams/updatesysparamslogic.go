package sysParams

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysParamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysParamsLogic {
	return &UpdateSysParamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysParamsLogic) UpdateSysParams(req *types.UpdateSysParamsReq) (resp *types.Response, err error) {
	// 1. 参数校验
	if req.ID == 0 {
		return nil, errors.New("参数ID不能为空")
	}

	// 2. 校验记录是否存在
	var exist model.SysParams
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&exist).Error; err != nil {
		return nil, errors.Wrap(err, "参数记录不存在")
	}

	// 3. 如果 key 变更了，校验唯一性
	if req.Key != "" && req.Key != exist.Key {
		var count int64
		if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysParams{}).Where("key = ? AND id != ?", req.Key, req.ID).Count(&count).Error; err != nil {
			return nil, errors.Wrap(err, "查询参数键是否重复失败")
		}
		if count > 0 {
			return nil, errors.New("参数键已存在")
		}
	}

	// 4. 更新记录
	updateMap := map[string]interface{}{
		"name":  req.Name,
		"key":   req.Key,
		"value": req.Value,
		"desc":  req.Desc,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysParams{}).Where("id = ?", req.ID).Updates(updateMap).Error; err != nil {
		l.Logger.Errorw("更新参数失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysParams"),
			logx.Field("action", "update"),
		)
		return nil, errors.Wrap(err, "更新参数失败")
	}

	return &types.Response{Msg: "更新成功"}, nil
}
