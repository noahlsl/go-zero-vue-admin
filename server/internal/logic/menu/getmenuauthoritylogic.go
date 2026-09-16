package menu

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.GetById) (resp *types.MenuRes, err error) {
	db := l.svcCtx.DB
	authorityId := uint(req.ID)

	// 查询角色关联的菜单ID列表
	var authMenus []model.SysAuthorityMenu
	if err = db.Where("sys_authority_authority_id = ?", authorityId).
		Find(&authMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色菜单关联失败")
	}

	menuIds := make([]string, 0, len(authMenus))
	for _, am := range authMenus {
		menuIds = append(menuIds, am.MenuId)
	}

	// 查询菜单详情
	var baseMenus []model.SysBaseMenu
	if err = db.Where("id in (?)", menuIds).Order("sort").
		Find(&baseMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单详情失败")
	}

	// 组装结果
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
			MenuId:      bm.ID,
		})
	}

	if menus == nil {
		menus = []types.SysMenu{}
	}

	return &types.MenuRes{
		AuthorityId: authorityId,
		Menus:       menus,
	}, nil
}
