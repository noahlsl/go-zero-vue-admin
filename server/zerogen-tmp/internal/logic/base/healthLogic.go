// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package base

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero/zerogen-tmp/internal/svc"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health() error {
	// todo: add your logic here and delete this line

	return nil
}
