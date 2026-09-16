// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package example

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
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
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	// 原 Gin 接口按操作者角色的数据权限范围过滤客户数据
	dataAuthorityIds, err := l.dataAuthorityIds()
	if err != nil {
		return nil, err
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.ExaCustomer{})
	db = db.Where("sys_user_authority_id in ?", dataAuthorityIds)

	var total int64
	if err = db.Count(&total).Error; err != nil {
		logx.Errorw("查询客户总数失败",
			logx.Field("error", err.Error()),
			logx.Field("module", "example"),
			logx.Field("action", "getCustomerList"),
		)
		return nil, errors.Wrap(err, "查询客户总数失败")
	}

	var customerList []model.ExaCustomer
	if err = db.Limit(limit).Offset(offset).Preload("SysUser").Find(&customerList).Error; err != nil {
		logx.Errorw("查询客户列表失败",
			logx.Field("error", err.Error()),
			logx.Field("module", "example"),
			logx.Field("action", "getCustomerList"),
		)
		return nil, errors.Wrap(err, "查询客户列表失败")
	}

	list := make([]types.ExaCustomer, 0, len(customerList))
	for _, c := range customerList {
		list = append(list, types.ExaCustomer{
			ID:                 c.ID,
			CreatedAt:          c.CreatedAt.Format(time.RFC3339),
			UpdatedAt:          c.UpdatedAt.Format(time.RFC3339),
			CustomerName:       c.CustomerName,
			CustomerPhoneData:  c.CustomerPhoneData,
			SysUserID:          c.SysUserID,
			SysUserAuthorityID: c.SysUserAuthorityID,
			SysUser:            conv.ToTypesUser(c.SysUser),
		})
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// dataAuthorityIds 查询当前操作者角色可访问的数据权限角色ID集合，
// 与原 Gin 后端 GetAuthorityInfo 预加载 DataAuthorityId 后取 AuthorityId 的行为一致。
func (l *GetCustomerListLogic) dataAuthorityIds() ([]uint, error) {
	authorityId := middleware.GetAuthorityId(l.ctx)
	if authorityId == 0 {
		return nil, errors.New("未获取到操作者角色信息")
	}

	var dataAuthorityIds []uint
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Table("sys_data_authority_id").
		Where("sys_authority_authority_id = ?", authorityId).
		Pluck("data_authority_id", &dataAuthorityIds).Error; err != nil {
		logx.Errorw("查询角色数据权限失败",
			logx.Field("error", err.Error()),
			logx.Field("authority_id", authorityId),
			logx.Field("module", "example"),
			logx.Field("action", "getCustomerList"),
		)
		return nil, errors.Wrap(err, "查询角色数据权限失败")
	}
	return dataAuthorityIds, nil
}
