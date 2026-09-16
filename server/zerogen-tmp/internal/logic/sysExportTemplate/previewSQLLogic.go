// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
