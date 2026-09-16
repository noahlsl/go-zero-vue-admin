// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package casbin

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCasbinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCasbinLogic {
	return &UpdateCasbinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCasbinLogic) UpdateCasbin(req *types.UpdateCasbinReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
