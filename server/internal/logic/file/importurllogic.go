package file

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportURLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportURLLogic {
	return &ImportURLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportURLLogic) ImportURL(req *types.ImportFileUrlReq) (resp *types.ExaFileRes, err error) {
	// TODO: 实现通过URL导入文件逻辑
	return nil, nil
}
