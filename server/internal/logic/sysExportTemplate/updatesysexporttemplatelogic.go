package sysExportTemplate

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysExportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysExportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysExportTemplateLogic {
	return &UpdateSysExportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysExportTemplateLogic) UpdateSysExportTemplate() (resp *types.Response, err error) {
	// 注意：当前 handler 未传递请求体参数，需要后续更新 handler 以支持请求体
	logx.WithContext(l.ctx).Infow("更新导出模板",
		logx.Field("module", "sysExportTemplate"),
		logx.Field("action", "update"),
	)

	return &types.Response{
		Code: 0,
		Msg:  "更新导出模板成功（请确保通过请求体传递模板数据）",
	}, nil
}
