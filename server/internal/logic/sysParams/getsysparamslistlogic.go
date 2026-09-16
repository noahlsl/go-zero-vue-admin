package sysParams

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysParamsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysParamsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysParamsListLogic {
	return &GetSysParamsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysParamsListLogic) GetSysParamsList(req *types.PageInfo) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 1. 计算分页参数
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysParams{})

	// 2. 条件筛选
	if req.Keyword != "" {
		db = db.Where("name LIKE ? OR `key` LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 3. 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询参数总数失败")
	}

	// 4. 分页查询
	var list []model.SysParams
	if err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error; err != nil {
		l.Logger.Errorw("查询参数列表失败",
			logx.Field("error", err),
			logx.Field("module", "sysParams"),
			logx.Field("action", "get_list"),
		)
		return nil, errors.Wrap(err, "查询参数列表失败")
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
