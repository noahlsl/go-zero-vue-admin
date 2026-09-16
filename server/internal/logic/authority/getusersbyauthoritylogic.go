package authority

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersByAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByAuthorityLogic {
	return &GetUsersByAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUsersByAuthority 获取拥有指定角色的用户ID列表
// 对齐原 Gin 实现：GetUserIdsByAuthorityId 仅返回 userIds，空结果返回空数组
func (l *GetUsersByAuthorityLogic) GetUsersByAuthority(req *types.SetRoleUsersReq) (resp []uint, err error) {
	// 1. 查询拥有该角色的所有用户ID
	var records []model.SysUserAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("sys_authority_authority_id = ?", req.AuthorityId).
		Find(&records).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色用户关联失败")
	}

	// 2. 收集用户ID，保持与原 Gin 一致：nil 时返回空数组
	userIds := make([]uint, 0, len(records))
	for _, r := range records {
		userIds = append(userIds, r.SysUserId)
	}

	return userIds, nil
}
