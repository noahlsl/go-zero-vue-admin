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
	customer := model.ExaCustomer{
		CustomerName:      req.CustomerName,
		CustomerPhoneData: req.CustomerPhoneData,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.ExaCustomer{}).Where("id = ?", req.ID).Updates(&customer).Error; err != nil {
		logx.Errorw("更新客户失败",
			logx.Field("error", err.Error()),
			logx.Field("id", req.ID),
			logx.Field("module", "example"),
			logx.Field("action", "updateCustomer"),
		)
		return nil, errors.Wrap(err, "更新客户失败")
	}
	return &types.Response{Msg: "更新成功"}, nil
}
