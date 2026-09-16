package api

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiRolesLogic {
	return &GetApiRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetApiRoles 获取拥有指定 API 权限的所有角色ID。
// 响应与原 Gin 一致：直接返回角色ID数组。
func (l *GetApiRolesLogic) GetApiRoles(req *types.GetApiRolesReq) (resp []uint, err error) {
	if req.Path == "" || req.Method == "" {
		return nil, errors.New("API路径和请求方法不能为空")
	}

	var casbinRules []model.SysCasbinRule
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("ptype = ? AND v1 = ? AND v2 = ?", "p", req.Path, req.Method).
		Find(&casbinRules).Error; err != nil {
		return nil, errors.Wrap(err, "查询API关联角色失败")
	}

	authorityIds := make([]uint, 0, len(casbinRules))
	for _, rule := range casbinRules {
		id, e := strconv.Atoi(rule.V0)
		if e == nil {
			authorityIds = append(authorityIds, uint(id))
		}
	}

	return authorityIds, nil
}
