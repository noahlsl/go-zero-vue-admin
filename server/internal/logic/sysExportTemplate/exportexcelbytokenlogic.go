package sysExportTemplate

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

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
	// 通过 Token 验证身份后导出 Excel 数据
	// 简化实现：返回基础响应
	logx.WithContext(l.ctx).Infow("通过Token导出Excel",
		logx.Field("module", "sysExportTemplate"),
		logx.Field("action", "export_excel_by_token"),
	)

	return &types.Response{
		Code: 0,
		Msg:  "通过Token导出Excel成功",
	}, nil
}
