// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package example

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FindCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCustomerLogic {
	return &FindCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCustomerLogic) FindCustomer(req *types.GetById) (resp *types.ExaCustomerRes, err error) {
	var customer model.ExaCustomer
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("客户不存在")
		}
		return nil, errors.Wrap(err, "查询客户详情失败")
	}
	resp = &types.ExaCustomerRes{
		Customer: types.ExaCustomer{
			ID:                 customer.ID,
			CreatedAt:          customer.CreatedAt.Format(time.RFC3339),
			UpdatedAt:          customer.UpdatedAt.Format(time.RFC3339),
			CustomerName:       customer.CustomerName,
			CustomerPhoneData:  customer.CustomerPhoneData,
			SysUserID:          customer.SysUserID,
			SysUserAuthorityID: customer.SysUserAuthorityID,
		},
	}
	return resp, nil
}
