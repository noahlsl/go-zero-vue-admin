package sysExportTemplate

import (
	"context"
	"fmt"
	"strings"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewSQLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPreviewSQLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewSQLLogic {
	return &PreviewSQLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PreviewSQLLogic) PreviewSQL(req *types.ExportTemplateReq) (resp *types.SqlPreviewRes, err error) {
	// 获取模板信息
	var template model.SysExportTemplate
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = ?", req.TemplateID).
		First(&template).Error; err != nil {
		return nil, errors.Wrap(err, "获取导出模板失败")
	}

	// 如果有自定义 SQL，直接返回
	if template.SQL != "" {
		return &types.SqlPreviewRes{Sql: template.SQL}, nil
	}

	// 构建 SQL 预览
	var sb strings.Builder
	sb.WriteString("SELECT ")
	sb.WriteString("*")
	sb.WriteString(" FROM ")
	sb.WriteString(template.Table)

	// 查询关联条件
	var conditions []model.SysExportTemplateCondition
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = ?", template.TemplateID).
		Find(&conditions).Error; err != nil {
		return nil, errors.Wrap(err, "查询模板条件失败")
	}

	if len(conditions) > 0 {
		var wheres []string
		for _, c := range conditions {
			wheres = append(wheres, fmt.Sprintf("%s %s ?", c.Column, c.Operator))
		}
		if len(wheres) > 0 {
			sb.WriteString(" WHERE ")
			sb.WriteString(strings.Join(wheres, " AND "))
		}
	}

	if template.Order != "" {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(template.Order)
	}

	if template.Limit != nil && *template.Limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", *template.Limit))
	}

	logx.WithContext(l.ctx).Infow("SQL预览成功",
		logx.Field("module", "sysExportTemplate"),
		logx.Field("action", "preview_sql"),
		logx.Field("template_id", template.TemplateID),
	)

	return &types.SqlPreviewRes{Sql: sb.String()}, nil
}
