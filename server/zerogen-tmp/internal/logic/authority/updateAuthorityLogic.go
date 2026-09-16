// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authority

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAuthorityLogic {
	return &UpdateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAuthorityLogic) UpdateAuthority(req *types.UpdateAuthorityReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
