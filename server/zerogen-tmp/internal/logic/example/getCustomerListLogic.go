// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package example

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCustomerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerListLogic {
	return &GetCustomerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerListLogic) GetCustomerList(req *types.GetExaCustomerListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
