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

type ExportExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportExcelLogic {
	return &ExportExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ExportExcel 导出表格，返回一次性导出链接。
// 原 Gin 通过 c.Query("templateID") 取值，响应 data 为形如
// /sysExportTemplate/exportExcelByToken?token=xxx 的字符串。
func (l *ExportExcelLogic) ExportExcel(req *types.ExportTemplateReq) (resp *types.Response, err error) {
	// 获取模板信息
	var template model.SysExportTemplate
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = ?", req.TemplateID).
		First(&template).Error; err != nil {
		return nil, errors.Wrap(err, "获取导出模板失败")
	}

	// TODO(未迁移): 一次性 token 缓存与 excelize 生成 xlsx 逻辑尚未实现，
	// 完整实现见 server/api/v1/system/sys_export_template.go ExportExcelByToken
	logx.WithContext(l.ctx).Infow("Excel导出准备完成",
		logx.Field("module", "sysExportTemplate"),
		logx.Field("action", "export_excel"),
		logx.Field("template_id", template.TemplateID),
	)

	return &types.Response{
		Code: 0,
		Data: fmt.Sprintf("/sysExportTemplate/exportExcelByToken?token=%s", template.TemplateID),
		Msg:  fmt.Sprintf("模板「%s」Excel导出准备完成", template.Name),
	}, nil
}

// buildSelectClause 构建 SELECT 子句
func buildSelectClause(templateInfo string) string {
	if templateInfo == "" {
		return "*"
	}
	// 从 templateInfo JSON 中提取列名
	// 简化处理：返回模板信息中的列
	return templateInfo
}

// buildJoinClause 构建 JOIN 子句
func buildJoinClause(joins []model.SysExportTemplateJoin) string {
	if len(joins) == 0 {
		return ""
	}
	var parts []string
	for _, j := range joins {
		parts = append(parts, fmt.Sprintf("%s %s ON %s", j.JOINS, j.Table, j.ON))
	}
	return strings.Join(parts, " ")
}
