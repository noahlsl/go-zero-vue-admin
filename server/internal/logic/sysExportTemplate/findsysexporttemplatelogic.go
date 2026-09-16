package sysExportTemplate

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FindSysExportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysExportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysExportTemplateLogic {
	return &FindSysExportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysExportTemplateLogic) FindSysExportTemplate(req *types.GetById) (resp *types.Response, err error) {
	var template model.SysExportTemplate
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("id = ?", req.ID).
		First(&template).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("导出模板不存在")
		}
		return nil, errors.Wrap(err, "查询导出模板失败")
	}

	// 查询关联的条件
	var conditions []model.SysExportTemplateCondition
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = ?", template.TemplateID).
		Find(&conditions).Error; err != nil {
		return nil, errors.Wrap(err, "查询模板条件失败")
	}

	// 查询关联的 Join 表
	var joins []model.SysExportTemplateJoin
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = ?", template.TemplateID).
		Find(&joins).Error; err != nil {
		return nil, errors.Wrap(err, "查询模板关联失败")
	}

	data := map[string]interface{}{
		"id":           template.ID,
		"dbName":       template.DBName,
		"name":         template.Name,
		"tableName":    template.Table,
		"templateID":   template.TemplateID,
		"templateInfo": template.TemplateInfo,
		"sql":          template.SQL,
		"importSQL":    template.ImportSQL,
		"limit":        template.Limit,
		"order":        template.Order,
		"conditions":   conditions,
		"joins":        joins,
	}

	return &types.Response{
		Code: 0,
		Data: data,
		Msg:  "获取成功",
	}, nil
}
