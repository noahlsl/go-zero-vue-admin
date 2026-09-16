// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetColumnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetColumnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetColumnLogic {
	return &GetColumnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetColumnLogic) GetColumn() (resp []types.ColumnInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
