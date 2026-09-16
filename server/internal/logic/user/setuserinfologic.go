package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserInfoLogic) SetUserInfo(req *types.ChangeUserInfoReq) (resp *types.Response, err error) {
	// 1. 如果指定了角色列表，先更新用户角色
	if len(req.AuthorityIds) > 0 {
		adminAuthorityId := middleware.GetAuthorityId(l.ctx)
		if err = l.setUserAuthorities(adminAuthorityId, req.ID, req.AuthorityIds); err != nil {
			return nil, errors.Wrap(err, "设置用户角色失败")
		}
	}

	// 2. 更新用户基本信息
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).
		Select("nick_name", "header_img", "phone", "email", "enable").
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"enable":     req.Enable,
		}).Error; err != nil {
		l.Logger.Errorw("设置用户信息失败",
			logx.Field("error", err),
			logx.Field("user_id", req.ID),
			logx.Field("module", "user"),
			logx.Field("action", "set_user_info"),
		)
		return nil, errors.Wrap(err, "设置用户信息失败")
	}

	return &types.Response{Msg: "设置成功"}, nil
}

// setUserAuthorities 批量设置用户角色
func (l *SetUserInfoLogic) setUserAuthorities(adminAuthorityId, userId uint, authorityIds []uint) error {
	return l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 查询用户是否存在
		var user model.SysUser
		if err := tx.Where("id = ?", userId).First(&user).Error; err != nil {
			return errors.Wrap(err, "查询用户数据失败")
		}

		// 删除用户原有的角色关联
		if err := tx.Where("sys_user_id = ?", userId).
			Delete(&model.SysUserAuthority{}).Error; err != nil {
			return errors.Wrap(err, "删除用户角色关联失败")
		}

		// 创建新的角色关联
		records := make([]model.SysUserAuthority, 0, len(authorityIds))
		for _, authorityId := range authorityIds {
			records = append(records, model.SysUserAuthority{
				SysUserId:               userId,
				SysAuthorityAuthorityId: authorityId,
			})
		}
		if err := tx.Create(&records).Error; err != nil {
			return errors.Wrap(err, "创建用户角色关联失败")
		}

		// 更新用户的默认角色
		if err := tx.Model(&user).Update("authority_id", authorityIds[0]).Error; err != nil {
			return errors.Wrap(err, "更新默认角色失败")
		}

		return nil
	})
}
