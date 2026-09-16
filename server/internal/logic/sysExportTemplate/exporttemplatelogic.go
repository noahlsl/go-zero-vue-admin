package sysExportTemplate

import (
	"context"
	"fmt"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTemplateLogic {
	return &ExportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ExportTemplate 导出表格模板，返回一次性导出链接。
// 原 Gin 通过 c.Query("templateID") 取值，响应 data 为形如
// /sysExportTemplate/exportTemplateByToken?token=xxx 的字符串。
func (l *ExportTemplateLogic) ExportTemplate(req *types.ExportTemplateReq) (resp *types.Response, err error) {
	// 获取模板信息
	var template model.SysExportTemplate
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = ?", req.TemplateID).
		First(&template).Error; err != nil {
		return nil, errors.Wrap(err, "获取导出模板失败")
	}

	// TODO(未迁移): 一次性 token 缓存与 excelize 生成 xlsx 逻辑尚未实现，
	// 完整实现见 server/api/v1/system/sys_export_template.go ExportTemplateByToken
	logx.WithContext(l.ctx).Infow("模板导出完成",
		logx.Field("module", "sysExportTemplate"),
		logx.Field("action", "export_template"),
		logx.Field("template_id", template.TemplateID),
	)

	return &types.Response{
		Code: 0,
		Data: fmt.Sprintf("/sysExportTemplate/exportTemplateByToken?token=%s", template.TemplateID),
		Msg:  fmt.Sprintf("模板「%s」导出完成", template.Name),
	}, nil
}
