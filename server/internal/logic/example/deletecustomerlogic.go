// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package example

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
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
	if err = l.svcCtx.DB.WithContext(l.ctx).Delete(&model.ExaCustomer{}, "id = ?", req.ID).Error; err != nil {
		logx.Errorw("删除客户失败",
			logx.Field("error", err.Error()),
			logx.Field("id", req.ID),
			logx.Field("module", "example"),
			logx.Field("action", "deleteCustomer"),
		)
		return nil, errors.Wrap(err, "删除客户失败")
	}
	return &types.Response{Msg: "删除成功"}, nil
}
