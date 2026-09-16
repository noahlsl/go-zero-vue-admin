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

type CreateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCustomerLogic {
	return &CreateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCustomerLogic) CreateCustomer(req *types.CreateExaCustomerReq) (resp *types.Response, err error) {
	customer := model.ExaCustomer{
		CustomerName:      req.CustomerName,
		CustomerPhoneData: req.CustomerPhoneData,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&customer).Error; err != nil {
		logx.Errorw("创建客户失败",
			logx.Field("error", err.Error()),
			logx.Field("module", "example"),
			logx.Field("action", "createCustomer"),
		)
		return nil, errors.Wrap(err, "创建客户失败")
	}
	return &types.Response{Msg: "创建成功"}, nil
}
