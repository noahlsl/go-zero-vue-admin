package sysExportTemplate

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysExportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysExportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysExportTemplateLogic {
	return &DeleteSysExportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysExportTemplateLogic) DeleteSysExportTemplate(req *types.GetById) (resp *types.Response, err error) {
	result := l.svcCtx.DB.WithContext(l.ctx).
		Delete(&model.SysExportTemplate{}, "id = ?", req.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "删除导出模板失败")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("导出模板不存在")
	}

	// 删除关联的条件和关联表数据
	if err := l.deleteRelatedData(req.ID); err != nil {
		return nil, err
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}

func (l *DeleteSysExportTemplateLogic) deleteRelatedData(id int) error {
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = (SELECT template_id FROM sys_export_templates WHERE id = ?)", id).
		Delete(&model.SysExportTemplateCondition{}).Error; err != nil {
		return errors.Wrap(err, "删除模板条件失败")
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("template_id = (SELECT template_id FROM sys_export_templates WHERE id = ?)", id).
		Delete(&model.SysExportTemplateJoin{}).Error; err != nil {
		return errors.Wrap(err, "删除模板关联失败")
	}

	return nil
}
