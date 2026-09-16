// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package api

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero/zerogen-tmp/internal/svc"
)

type EnterSyncApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnterSyncApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnterSyncApiLogic {
	return &EnterSyncApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnterSyncApiLogic) EnterSyncApi() error {
	// todo: add your logic here and delete this line

	return nil
}
