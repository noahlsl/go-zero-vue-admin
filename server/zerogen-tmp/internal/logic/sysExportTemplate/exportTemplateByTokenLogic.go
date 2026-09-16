// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTemplateByTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportTemplateByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTemplateByTokenLogic {
	return &ExportTemplateByTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportTemplateByTokenLogic) ExportTemplateByToken() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
