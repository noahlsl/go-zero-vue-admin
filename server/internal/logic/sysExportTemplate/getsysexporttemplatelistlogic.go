package sysExportTemplate

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysExportTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysExportTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysExportTemplateListLogic {
	return &GetSysExportTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysExportTemplateListLogic) GetSysExportTemplateList(req *types.PageInfo) (resp *types.PageResult, err error) {
	var total int64
	var templates []model.SysExportTemplate

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysExportTemplate{})

	if req.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "统计导出模板失败")
	}

	offset, limit := templatePageConvert(req.Page, req.PageSize)
	if err := db.Offset(offset).Limit(limit).
		Order("id DESC").
		Find(&templates).Error; err != nil {
		return nil, errors.Wrap(err, "查询导出模板列表失败")
	}

	list := make([]map[string]interface{}, 0, len(templates))
	for _, t := range templates {
		list = append(list, map[string]interface{}{
			"id":           t.ID,
			"dbName":       t.DBName,
			"name":         t.Name,
			"tableName":    t.Table,
			"templateID":   t.TemplateID,
			"templateInfo": t.TemplateInfo,
			"sql":          t.SQL,
			"importSQL":    t.ImportSQL,
			"limit":        t.Limit,
			"order":        t.Order,
		})
	}

	resp = &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return resp, nil
}

func templatePageConvert(page, pageSize int) (offset, limit int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return (page - 1) * pageSize, pageSize
}
