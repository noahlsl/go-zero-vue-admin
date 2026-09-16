// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportSysDictionaryLogic {
	return &ExportSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportSysDictionaryLogic) ExportSysDictionary(req *types.GetById) (resp *types.ExportSysDictionaryRes, err error) {
	// todo: add your logic here and delete this line

	return
}
