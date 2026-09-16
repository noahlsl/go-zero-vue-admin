// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package example

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomerLogic {
	return &DeleteCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCustomerLogic) DeleteCustomer(req *types.GetById) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
