package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// allowedOrderKeys 允许排序的字段白名单
var allowedOrderKeys = map[string]bool{
	"id":        true,
	"username":  true,
	"nick_name": true,
	"phone":     true,
	"email":     true,
}

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 1. 计算分页参数
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{})

	// 2. 条件筛选
	db = buildFilterQuery(db, req)

	// 3. 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询用户总数失败")
	}

	// 4. 排序
	orderStr := buildOrderStr(req)
	db = db.Order(orderStr)

	// 5. 分页查询
	var userList []model.SysUser
	if err = db.Limit(limit).Offset(offset).
		Preload("Authorities").Preload("Authority").
		Find(&userList).Error; err != nil {
		l.Logger.Errorw("查询用户列表失败",
			logx.Field("error", err),
			logx.Field("module", "user"),
			logx.Field("action", "get_user_list"),
		)
		return nil, errors.Wrap(err, "查询用户列表失败")
	}

	// 6. 转换为响应类型
	list := make([]types.SysUser, 0, len(userList))
	for _, u := range userList {
		list = append(list, conv.ToTypesUser(u))
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// buildFilterQuery 构建筛选条件
func buildFilterQuery(db *gorm.DB, req *types.GetUserListReq) *gorm.DB {
	if req.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+req.NickName+"%")
	}
	if req.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+req.Phone+"%")
	}
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Email != "" {
		db = db.Where("email LIKE ?", "%"+req.Email+"%")
	}
	return db
}

// buildOrderStr 构建排序字符串
func buildOrderStr(req *types.GetUserListReq) string {
	if req.OrderKey != "" && allowedOrderKeys[req.OrderKey] {
		orderStr := req.OrderKey
		if req.Desc {
			orderStr += " desc"
		}
		return orderStr
	}
	return "id desc"
}
