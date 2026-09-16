package sysExportTemplate

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysExportTemplateByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysExportTemplateByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysExportTemplateByIdsLogic {
	return &DeleteSysExportTemplateByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysExportTemplateByIdsLogic) DeleteSysExportTemplateByIds(req *types.IdsReq) (resp *types.Response, err error) {
	if len(req.Ids) == 0 {
		return nil, errors.New("ID列表不能为空")
	}

	// 查询关联的 template_id
	var templateIDs []string
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysExportTemplate{}).
		Where("id IN ?", req.Ids).
		Pluck("template_id", &templateIDs).Error; err != nil {
		return nil, errors.Wrap(err, "查询模板标识失败")
	}

	// 删除模板记录
	result := l.svcCtx.DB.WithContext(l.ctx).
		Where("id IN ?", req.Ids).
		Delete(&model.SysExportTemplate{})
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "批量删除导出模板失败")
	}

	// 删除关联数据
	if len(templateIDs) > 0 {
		if err := l.svcCtx.DB.WithContext(l.ctx).
			Where("template_id IN ?", templateIDs).
			Delete(&model.SysExportTemplateCondition{}).Error; err != nil {
			logx.WithContext(l.ctx).Errorw("删除模板条件失败",
				logx.Field("error", err.Error()),
				logx.Field("module", "sysExportTemplate"),
				logx.Field("action", "batch_delete_related"),
			)
		}
		if err := l.svcCtx.DB.WithContext(l.ctx).
			Where("template_id IN ?", templateIDs).
			Delete(&model.SysExportTemplateJoin{}).Error; err != nil {
			logx.WithContext(l.ctx).Errorw("删除模板关联失败",
				logx.Field("error", err.Error()),
				logx.Field("module", "sysExportTemplate"),
				logx.Field("action", "batch_delete_related"),
			)
		}
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
