package authority

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SetRoleUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleUsersLogic {
	return &SetRoleUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleUsersLogic) SetRoleUsers(req *types.SetRoleUsersReq) (resp *types.Response, err error) {
	// 1. 参数校验
	if req.AuthorityId == 0 {
		return nil, errors.New("角色ID不能为空")
	}

	// 2. 事务：全量覆盖角色关联的用户列表
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 2.1 查询当前拥有该角色的所有用户关联
		var existingRecords []model.SysUserAuthority
		if e := tx.Where("sys_authority_authority_id = ?", req.AuthorityId).
			Find(&existingRecords).Error; e != nil {
			return errors.Wrap(e, "查询当前角色用户关联失败")
		}

		currentSet := make(map[uint]struct{})
		for _, r := range existingRecords {
			currentSet[r.SysUserId] = struct{}{}
		}

		targetSet := make(map[uint]struct{})
		for _, id := range req.UserIds {
			targetSet[id] = struct{}{}
		}

		// 2.2 删除该角色所有已有的用户关联
		if e := tx.Delete(&model.SysUserAuthority{},
			"sys_authority_authority_id = ?", req.AuthorityId).Error; e != nil {
			return errors.Wrap(e, "删除角色用户关联失败")
		}

		// 2.3 对被移除的用户：若该角色是其主角色，则将主角色切换为其剩余的其他角色
		for userId := range currentSet {
			if _, ok := targetSet[userId]; ok {
				continue // 仍在目标列表中，不处理
			}
			var user model.SysUser
			if e := tx.First(&user, "id = ?", userId).Error; e != nil {
				continue
			}
			if user.AuthorityId == req.AuthorityId {
				// 从剩余关联中找另一个角色作为主角色
				var another model.SysUserAuthority
				if e := tx.Where("sys_user_id = ?", userId).First(&another).Error; e != nil {
					// 没有其他角色，跳过
					continue
				}
				if e := tx.Model(&model.SysUser{}).Where("id = ?", userId).
					Update("authority_id", another.SysAuthorityAuthorityId).Error; e != nil {
					return errors.Wrap(e, "更新用户主角色失败")
				}
			}
		}

		// 2.4 批量插入新的关联记录
		if len(req.UserIds) > 0 {
			newRecords := make([]model.SysUserAuthority, 0, len(req.UserIds))
			for _, userId := range req.UserIds {
				newRecords = append(newRecords, model.SysUserAuthority{
					SysUserId:               userId,
					SysAuthorityAuthorityId: req.AuthorityId,
				})
			}
			if e := tx.Create(&newRecords).Error; e != nil {
				return errors.Wrap(e, "批量插入角色用户关联失败")
			}
		}

		return nil
	})
	if err != nil {
		l.Logger.Errorw("设置角色用户失败",
			logx.Field("error", err),
			logx.Field("authority_id", req.AuthorityId),
			logx.Field("module", "authority"),
			logx.Field("action", "set_role_users"),
		)
		return nil, err
	}

	resp = &types.Response{
		Code: 0,
		Msg:  "设置成功",
	}
	return resp, nil
}
