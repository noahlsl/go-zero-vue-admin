package menu

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.GetById) (resp []types.SysMenu, err error) {
	db := l.svcCtx.DB
	authorityId := uint(req.ID)

	// 加载该角色关联的所有菜单ID
	var authMenus []model.SysAuthorityMenu
	if err = db.Where("sys_authority_authority_id = ?", authorityId).
		Find(&authMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色菜单关联失败")
	}

	if len(authMenus) == 0 {
		return []types.SysMenu{}, nil
	}

	menuIds := make([]string, 0, len(authMenus))
	for _, am := range authMenus {
		menuIds = append(menuIds, am.MenuId)
	}

	// 加载菜单基本信息
	var baseMenus []model.SysBaseMenu
	if err = db.Where("id in (?)", menuIds).Order("sort").Find(&baseMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单列表失败")
	}

	// 组装 SysMenu 列表
	menus := make([]types.SysMenu, 0, len(baseMenus))
	for _, bm := range baseMenus {
		menus = append(menus, types.SysMenu{
			ID:        bm.ID,
			MenuLevel: bm.MenuLevel,
			ParentId:  bm.ParentId,
			Path:      bm.Path,
			Name:      bm.Name,
			Component: bm.Component,
			Sort:      bm.Sort,
			Meta: types.Meta{
				KeepAlive:   bm.KeepAlive,
				DefaultMenu: bm.DefaultMenu,
				Title:       bm.Title,
				Icon:        bm.Icon,
				CloseTab:    bm.CloseTab,
			},
			Hidden:      bm.Hidden,
			AuthorityId: authorityId,
		})
	}

	// 加载角色按钮权限
	if err = l.loadAuthorityBtns(menus, authorityId); err != nil {
		return nil, err
	}

	return buildMenuTree(menus), nil
}

func (l *GetMenuListLogic) loadAuthorityBtns(menus []types.SysMenu, authorityId uint) error {
	db := l.svcCtx.DB

	var authBtns []model.SysAuthorityBtn
	if err := db.Where("authority_id = ?", authorityId).Find(&authBtns).Error; err != nil {
		return errors.Wrap(err, "查询角色按钮权限失败")
	}

	if len(authBtns) == 0 {
		return nil
	}

	// 收集按钮ID列表
	btnIds := make([]uint, 0, len(authBtns))
	for _, ab := range authBtns {
		btnIds = append(btnIds, ab.SysBaseMenuBtnID)
	}

	// 查询基础菜单按钮
	var baseBtns []model.SysBaseMenuBtn
	if err := db.Where("id IN (?)", btnIds).Find(&baseBtns).Error; err != nil {
		return errors.Wrap(err, "查询基础菜单按钮失败")
	}

	// 构建按钮ID到名称的映射
	btnNameMap := make(map[uint]string)
	for _, bb := range baseBtns {
		btnNameMap[bb.ID] = bb.Name
	}

	// 构建菜单ID到按钮列表的映射
	btnMap := make(map[uint][]types.SysAuthorityBtn)
	for _, ab := range authBtns {
		name := btnNameMap[ab.SysBaseMenuBtnID]
		btnMap[ab.SysMenuID] = append(btnMap[ab.SysMenuID], types.SysAuthorityBtn{
			ID:   ab.ID,
			Name: name,
			Desc: "",
		})
	}

	for i := range menus {
		if btnList, ok := btnMap[menus[i].ID]; ok {
			menus[i].Buttons = btnList
		} else {
			menus[i].Buttons = []types.SysAuthorityBtn{}
		}
	}

	return nil
}
