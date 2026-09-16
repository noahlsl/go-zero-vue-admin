// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package base

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero/zerogen-tmp/internal/svc"
)

type InitDBLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDBLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDBLogic {
	return &InitDBLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDBLogic) InitDB() error {
	// todo: add your logic here and delete this line

	return nil
}
