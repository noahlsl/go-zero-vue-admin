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

type SetUserAuthoritiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserAuthoritiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserAuthoritiesLogic {
	return &SetUserAuthoritiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserAuthoritiesLogic) SetUserAuthorities(req *types.SetUserAuthoritiesReq) (resp *types.Response, err error) {
	// 1. 从 context 获取当前操作者的角色 ID
	adminAuthorityId := middleware.GetAuthorityId(l.ctx)
	if adminAuthorityId == 0 {
		return nil, errors.New("未获取到操作者角色信息")
	}

	// 2. 事务执行批量设置
	if err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 查询用户是否存在
		var user model.SysUser
		if err := tx.Where("id = ?", req.ID).First(&user).Error; err != nil {
			return errors.Wrap(err, "查询用户数据失败")
		}

		// 删除用户原有的角色关联
		if err := tx.Where("sys_user_id = ?", req.ID).
			Delete(&model.SysUserAuthority{}).Error; err != nil {
			return errors.Wrap(err, "删除用户角色关联失败")
		}

		// 创建新的角色关联
		records := make([]model.SysUserAuthority, 0, len(req.AuthorityIds))
		for _, authorityId := range req.AuthorityIds {
			records = append(records, model.SysUserAuthority{
				SysUserId:               req.ID,
				SysAuthorityAuthorityId: authorityId,
			})
		}
		if err := tx.Create(&records).Error; err != nil {
			return errors.Wrap(err, "创建用户角色关联失败")
		}

		// 更新用户的默认角色
		if err := tx.Model(&user).Update("authority_id", req.AuthorityIds[0]).Error; err != nil {
			return errors.Wrap(err, "更新默认角色失败")
		}

		return nil
	}); err != nil {
		l.Logger.Errorw("批量设置用户角色失败",
			logx.Field("error", err),
			logx.Field("user_id", req.ID),
			logx.Field("admin_authority_id", adminAuthorityId),
			logx.Field("module", "user"),
			logx.Field("action", "set_user_authorities"),
		)
		return nil, err
	}

	return &types.Response{Msg: "修改成功"}, nil
}
