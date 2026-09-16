// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysExportTemplate

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysExportTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysExportTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysExportTemplateListLogic {
	return &GetSysExportTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysExportTemplateListLogic) GetSysExportTemplateList(req *types.PageInfo) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
