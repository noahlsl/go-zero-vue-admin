package api

import (
	"context"
	"fmt"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.GetApiListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysApi{})

	// 条件过滤
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Description != "" {
		db = db.Where("description LIKE ?", "%"+req.Description+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.ApiGroup != "" {
		db = db.Where("api_group = ?", req.ApiGroup)
	}

	// 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "统计API总数失败")
	}

	// 排序
	orderStr := "id desc"
	if req.OrderKey != "" {
		allowedOrders := map[string]bool{
			"id": true, "path": true, "api_group": true,
			"description": true, "method": true,
		}
		if !allowedOrders[req.OrderKey] {
			return nil, fmt.Errorf("非法的排序字段: %v", req.OrderKey)
		}
		orderStr = req.OrderKey
		if req.Desc {
			orderStr = req.OrderKey + " desc"
		}
	}

	// 分页查询
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var apiList []model.SysApi
	if err = db.Limit(limit).Offset(offset).Order(orderStr).Find(&apiList).Error; err != nil {
		return nil, errors.Wrap(err, "查询API列表失败")
	}

	// 转换为 types.SysApi
	list := make([]types.SysApi, 0, len(apiList))
	for _, item := range apiList {
		list = append(list, types.SysApi{
			ID:          item.ID,
			CreatedAt:   item.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
			Method:      item.Method,
			Path:        item.Path,
			ApiGroup:    item.ApiGroup,
			Description: item.Description,
		})
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
