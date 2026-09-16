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

type AddBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBaseMenuLogic {
	return &AddBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBaseMenuLogic) AddBaseMenu(req *types.AddMenuReq) (resp *types.Response, err error) {
	db := l.svcCtx.DB

	// 检查 name 是否重复
	var count int64
	if err = db.Model(&model.SysBaseMenu{}).Where("name = ?", req.Name).Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单名称失败")
	}
	if count > 0 {
		return nil, errors.New("存在重复 name，请修改 name")
	}

	// 校验父菜单
	if req.ParentId != 0 {
		if err = l.validateParentMenu(req.ParentId); err != nil {
			return nil, err
		}
	}

	// 事务创建菜单及相关子表
	err = db.Transaction(func(tx *gorm.DB) error {
		menu := model.SysBaseMenu{
			MenuLevel:   req.MenuLevel,
			ParentId:    req.ParentId,
			Path:        req.Path,
			Name:        req.Name,
			Component:   req.Component,
			Sort:        req.Sort,
			ActiveName:  "",
			KeepAlive:   req.KeepAlive,
			DefaultMenu: req.DefaultMenu,
			Title:       req.Title,
			Icon:        req.Icon,
			CloseTab:    req.CloseTab,
			Hidden:      req.Hidden,
		}
		if err := tx.Create(&menu).Error; err != nil {
			return errors.Wrap(err, "创建菜单失败")
		}

		// 创建菜单按钮
		if len(req.Button) > 0 {
			btns := make([]model.SysBaseMenuBtn, 0, len(req.Button))
			for _, b := range req.Button {
				btns = append(btns, model.SysBaseMenuBtn{
					Name:          b.Name,
					Desc:          b.Desc,
					SysBaseMenuID: menu.ID,
				})
			}
			if err := tx.Create(&btns).Error; err != nil {
				return errors.Wrap(err, "创建菜单按钮失败")
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "添加菜单事务失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "添加成功",
	}, nil
}

func (l *AddBaseMenuLogic) validateParentMenu(parentId uint) error {
	db := l.svcCtx.DB

	var parentMenu model.SysBaseMenu
	if err := db.First(&parentMenu, parentId).Error; err != nil {
		return errors.Wrap(err, "父菜单不存在")
	}

	// 检查父菜单下现有子菜单数量
	var childCount int64
	if err := db.Model(&model.SysBaseMenu{}).Where("parent_id = ?", parentId).Count(&childCount).Error; err != nil {
		return errors.Wrap(err, "查询子菜单数量失败")
	}

	// 如果父菜单原本是叶子菜单，现在变成枝干菜单，需要清空其权限分配
	if childCount == 0 {
		var defaultRouterCount int64
		if err := db.Model(&model.SysAuthority{}).
			Where("default_router = ?", parentMenu.Name).
			Count(&defaultRouterCount).Error; err != nil {
			return errors.Wrap(err, "查询默认路由引用失败")
		}
		if defaultRouterCount > 0 {
			return errors.New("父菜单已被其他角色的首页占用，请先释放父菜单的首页权限")
		}

		// 清空父菜单的所有权限分配
		if err := db.Where("sys_base_menu_id = ?", parentId).
			Delete(&model.SysAuthorityMenu{}).Error; err != nil {
			return errors.Wrap(err, "清空父菜单权限分配失败")
		}
	}

	return nil
}
