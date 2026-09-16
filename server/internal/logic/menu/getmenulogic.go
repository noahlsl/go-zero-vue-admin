package menu

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetMenu 获取当前登录用户角色对应的动态路由菜单树。
// 原 Gin 接口无入参，前端不传请求体，authorityId 从 JWT 上下文获取。
func (l *GetMenuLogic) GetMenu() (resp *types.MenuRes, err error) {
	db := l.svcCtx.DB
	authorityId := middleware.GetAuthorityId(l.ctx)

	// 加载该角色关联的所有菜单ID
	var authMenus []model.SysAuthorityMenu
	if err = db.Where("sys_authority_authority_id = ?", authorityId).
		Find(&authMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色菜单关联失败")
	}

	if len(authMenus) == 0 {
		return &types.MenuRes{
			AuthorityId: authorityId,
			Menus:       []types.SysMenu{},
		}, nil
	}

	menuIds := make([]string, 0, len(authMenus))
	for _, am := range authMenus {
		menuIds = append(menuIds, am.MenuId)
	}

	// 加载菜单信息和参数
	var baseMenus []model.SysBaseMenu
	if err = db.Where("id in (?)", menuIds).Order("sort").
		Preload("Parameters").Find(&baseMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单列表失败")
	}

	// 组装 SysMenu
	menus := make([]types.SysMenu, 0, len(baseMenus))
	for _, bm := range baseMenus {
		params := make([]types.SysMenuParameter, 0, len(bm.Parameters))
		for _, p := range bm.Parameters {
			params = append(params, types.SysMenuParameter{
				ID:            p.ID,
				SysBaseMenuID: p.SysBaseMenuID,
				Type:          p.Type,
				Key:           p.Key,
				Value:         p.Value,
			})
		}
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
			MenuId:      bm.ID,
			Parameters:  params,
		})
	}

	// 加载角色按钮权限
	if err = l.loadAuthorityBtns(menus, authorityId); err != nil {
		return nil, err
	}

	tree := buildMenuTree(menus)
	return &types.MenuRes{
		AuthorityId: authorityId,
		Menus:       tree,
	}, nil
}

func (l *GetMenuLogic) loadAuthorityBtns(menus []types.SysMenu, authorityId uint) error {
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
