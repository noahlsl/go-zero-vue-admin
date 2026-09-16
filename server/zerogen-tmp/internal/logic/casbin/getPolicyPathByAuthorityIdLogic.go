// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package casbin

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyPathByAuthorityIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyPathByAuthorityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyPathByAuthorityIdLogic {
	return &GetPolicyPathByAuthorityIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPolicyPathByAuthorityIdLogic) GetPolicyPathByAuthorityId(req *types.GetPolicyPathByAuthorityIdReq) (resp *types.CasbinRes, err error) {
	// todo: add your logic here and delete this line

	return
}
