package menu

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBaseMenuLogic {
	return &DeleteBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBaseMenuLogic) DeleteBaseMenu(req *types.DeleteMenuReq) (resp *types.Response, err error) {
	db := l.svcCtx.DB

	// 检查是否存在子菜单
	var childCount int64
	if err = db.Model(&model.SysBaseMenu{}).Where("parent_id = ?", req.ID).
		Count(&childCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询子菜单失败")
	}
	if childCount > 0 {
		return nil, errors.New("此菜单存在子菜单不可删除")
	}

	// 校验菜单是否存在
	var menu model.SysBaseMenu
	if err = db.First(&menu, req.ID).Error; err != nil {
		return nil, errors.Wrap(err, "记录不存在")
	}

	// 检查是否有角色将该菜单作为首页
	var defaultRouterCount int64
	if err = db.Model(&model.SysAuthority{}).
		Where("default_router = ?", menu.Name).
		Count(&defaultRouterCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询默认路由引用失败")
	}
	if defaultRouterCount > 0 {
		return nil, errors.New("此菜单有角色正在作为首页，不可删除")
	}

	// 事务删除
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.SysBaseMenu{}, "id = ?", req.ID).Error; err != nil {
			return errors.Wrap(err, "删除菜单失败")
		}
		if err := tx.Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", req.ID).Error; err != nil {
			return errors.Wrap(err, "删除菜单参数失败")
		}
		if err := tx.Delete(&model.SysBaseMenuBtn{}, "sys_base_menu_id = ?", req.ID).Error; err != nil {
			return errors.Wrap(err, "删除菜单按钮失败")
		}
		if err := tx.Delete(&model.SysAuthorityBtn{}, "sys_menu_id = ?", req.ID).Error; err != nil {
			return errors.Wrap(err, "删除角色按钮权限失败")
		}
		if err := tx.Delete(&model.SysAuthorityMenu{}, "sys_base_menu_id = ?", req.ID).Error; err != nil {
			return errors.Wrap(err, "删除角色菜单关联失败")
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "删除菜单事务失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
