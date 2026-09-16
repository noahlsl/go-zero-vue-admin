package authority

import (
	"context"
	"fmt"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAuthorityLogic {
	return &DeleteAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAuthorityLogic) DeleteAuthority(req *types.GetById) (resp *types.Response, err error) {
	authorityId := uint(req.ID)

	// 1. 校验角色是否存在
	var authority model.SysAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", authorityId).First(&authority).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("该角色不存在")
		}
		return nil, errors.Wrap(err, "查询角色失败")
	}

	// 2. 检查是否有用户正在使用此角色
	var userCount int64
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).Where("authority_id = ?", authorityId).Count(&userCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色用户数量失败")
	}
	if userCount > 0 {
		return nil, errors.New("此角色有用户正在使用禁止删除")
	}

	// 3. 检查是否有关联用户（sys_user_authority 表）
	var relCount int64
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUserAuthority{}).
		Where("sys_authority_authority_id = ?", authorityId).Count(&relCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色关联用户失败")
	}
	if relCount > 0 {
		return nil, errors.New("此角色有用户正在使用禁止删除")
	}

	// 4. 检查是否有子角色
	var childCount int64
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysAuthority{}).
		Where("parent_id = ?", authorityId).Count(&childCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询子角色失败")
	}
	if childCount > 0 {
		return nil, errors.New("此角色存在子角色不允许删除")
	}

	// 5. 事务删除角色及相关关联数据
	authorityIdStr := fmt.Sprintf("%d", authorityId)
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 删除角色
		if e := tx.Unscoped().Delete(&authority).Error; e != nil {
			return errors.Wrap(e, "删除角色失败")
		}

		// 删除菜单关联
		if e := tx.Exec("DELETE FROM sys_authority_menus WHERE sys_authority_authority_id = ?",
			authorityIdStr).Error; e != nil {
			return errors.Wrap(e, "删除角色菜单关联失败")
		}

		// 删除用户角色关联
		if e := tx.Delete(&model.SysUserAuthority{},
			"sys_authority_authority_id = ?", authorityId).Error; e != nil {
			return errors.Wrap(e, "删除用户角色关联失败")
		}

		// 删除角色按钮权限
		if e := tx.Where("authority_id = ?", authorityId).
			Delete(&[]model.SysAuthorityBtn{}).Error; e != nil {
			return errors.Wrap(e, "删除角色按钮权限失败")
		}

		return nil
	})
	if err != nil {
		l.Logger.Errorw("删除角色失败",
			logx.Field("error", err),
			logx.Field("authority_id", authorityId),
			logx.Field("module", "authority"),
			logx.Field("action", "delete"),
		)
		return nil, err
	}

	resp = &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}
	return resp, nil
}
