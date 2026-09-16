// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero/zerogen-tmp/internal/svc"
)

type IgnoreApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIgnoreApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IgnoreApiLogic {
	return &IgnoreApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IgnoreApiLogic) IgnoreApi() error {
	// todo: add your logic here and delete this line

	return nil
}
