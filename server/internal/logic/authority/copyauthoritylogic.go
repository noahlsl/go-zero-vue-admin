package authority

import (
	"context"
	"fmt"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CopyAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyAuthorityLogic {
	return &CopyAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyAuthorityLogic) CopyAuthority(req *types.CopyAuthorityReq) (resp *types.Response, err error) {
	// 1. 校验新角色ID是否已存在
	var existAuth model.SysAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", req.AuthorityId).First(&existAuth).Error; err == nil {
		return nil, errors.New("存在相同角色id")
	} else if err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "查询角色是否已存在失败")
	}

	// 2. 查询旧角色关联的菜单
	var oldMenuIds []string
	if err = l.svcCtx.DB.WithContext(l.ctx).Table("sys_authority_menus").
		Where("sys_authority_authority_id = ?", req.OldAuthorityId).
		Pluck("sys_base_menu_id", &oldMenuIds).Error; err != nil {
		return nil, errors.Wrap(err, "查询旧角色菜单关联失败")
	}

	// 3. 事务：创建新角色 + 关联菜单 + 复制按钮权限
	newAuth := model.SysAuthority{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      &req.ParentId,
		DefaultRouter: req.DefaultRouter,
		DataScope:     req.DataScope,
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 创建新角色
		if e := tx.Create(&newAuth).Error; e != nil {
			return errors.Wrap(e, "创建复制角色失败")
		}

		// 关联菜单
		if len(oldMenuIds) > 0 {
			menuRecords := make([]model.SysAuthorityMenu, 0, len(oldMenuIds))
			for _, menuId := range oldMenuIds {
				menuRecords = append(menuRecords, model.SysAuthorityMenu{
					MenuId:      menuId,
					AuthorityId: l.uintToString(req.AuthorityId),
				})
			}
			if e := tx.Create(&menuRecords).Error; e != nil {
				return errors.Wrap(e, "复制角色菜单关联失败")
			}
		}

		// 复制按钮权限
		var btns []model.SysAuthorityBtn
		if e := tx.Where("authority_id = ?", req.OldAuthorityId).Find(&btns).Error; e != nil {
			return errors.Wrap(e, "查询旧角色按钮权限失败")
		}
		if len(btns) > 0 {
			for i := range btns {
				btns[i].ID = 0
				btns[i].AuthorityId = req.AuthorityId
			}
			if e := tx.Create(&btns).Error; e != nil {
				return errors.Wrap(e, "复制角色按钮权限失败")
			}
		}

		return nil
	})
	if err != nil {
		l.Logger.Errorw("复制角色失败",
			logx.Field("error", err),
			logx.Field("old_authority_id", req.OldAuthorityId),
			logx.Field("new_authority_id", req.AuthorityId),
			logx.Field("module", "authority"),
			logx.Field("action", "copy"),
		)
		return nil, err
	}

	resp = &types.Response{
		Code: 0,
		Msg:  "拷贝成功",
		Data: conv.ToTypesAuthority(newAuth),
	}
	return resp, nil
}

func (l *CopyAuthorityLogic) uintToString(v uint) string {
	return fmt.Sprintf("%d", v)
}
