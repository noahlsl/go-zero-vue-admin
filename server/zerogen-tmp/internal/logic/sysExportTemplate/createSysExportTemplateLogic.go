// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysExportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysExportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysExportTemplateLogic {
	return &CreateSysExportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysExportTemplateLogic) CreateSysExportTemplate() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
