// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package file

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportFileUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportFileUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportFileUrlLogic {
	return &ImportFileUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportFileUrlLogic) ImportFileUrl(req *types.ImportFileUrlReq) (resp *types.ExaFileRes, err error) {
	// todo: add your logic here and delete this line

	return
}
