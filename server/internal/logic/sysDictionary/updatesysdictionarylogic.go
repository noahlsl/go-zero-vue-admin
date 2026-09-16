package sysDictionary

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictionaryLogic {
	return &UpdateSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysDictionaryLogic) UpdateSysDictionary(req *types.UpdateSysDictionaryReq) (resp *types.Response, err error) {
	// 查询原字典
	var dict model.SysDictionary
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&dict).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典数据失败")
	}

	// 如果 type 发生变更，检查新 type 是否已存在
	if dict.Type != req.Type {
		var count int64
		if err = l.svcCtx.DB.WithContext(l.ctx).
			Model(&model.SysDictionary{}).
			Where("type = ?", req.Type).
			Count(&count).Error; err != nil {
			return nil, errors.Wrap(err, "查询字典类型是否存在失败")
		}
		if count > 0 {
			return nil, errors.New("存在相同的type，不允许创建")
		}
	}

	updates := map[string]interface{}{
		"name":      req.Name,
		"type":      req.Type,
		"status":    req.Status,
		"desc":      req.Desc,
		"parent_id": req.ParentID,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&dict).Updates(updates).Error; err != nil {
		return nil, errors.Wrap(err, "更新字典失败")
	}

	return &types.Response{Code: 0, Msg: "修改成功"}, nil
}
