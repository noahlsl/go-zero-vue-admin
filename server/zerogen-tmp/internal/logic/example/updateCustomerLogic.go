// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package example

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerLogic) UpdateCustomer(req *types.UpdateExaCustomerReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
