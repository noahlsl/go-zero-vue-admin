// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportExcelByTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportExcelByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportExcelByTokenLogic {
	return &ExportExcelByTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportExcelByTokenLogic) ExportExcelByToken() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
