// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package system

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSystemConfigLogic {
	return &SetSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSystemConfigLogic) SetSystemConfig(req *types.SystemConfig) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
