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

type DeleteCustomerByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCustomerByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomerByIdsLogic {
	return &DeleteCustomerByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCustomerByIdsLogic) DeleteCustomerByIds(req *types.IdsReq) (resp *types.Response, err error) {
	if err = l.svcCtx.DB.WithContext(l.ctx).Delete(&model.ExaCustomer{}, "id in ?", req.Ids).Error; err != nil {
		logx.Errorw("批量删除客户失败",
			logx.Field("error", err.Error()),
			logx.Field("ids", req.Ids),
			logx.Field("module", "example"),
			logx.Field("action", "deleteCustomerByIds"),
		)
		return nil, errors.Wrap(err, "批量删除客户失败")
	}
	return &types.Response{}, nil
}
