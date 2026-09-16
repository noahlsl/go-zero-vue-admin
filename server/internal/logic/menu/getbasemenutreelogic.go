package menu

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseMenuTreeLogic {
	return &GetBaseMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBaseMenuTreeLogic) GetBaseMenuTree() (resp *types.MenuTreeRes, err error) {
	db := l.svcCtx.DB

	var baseMenus []model.SysBaseMenu
	if err = db.Order("sort").Preload("Parameters").Preload("MenuBtn").
		Find(&baseMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单树失败")
	}

	menus := make([]types.SysMenu, 0, len(baseMenus))
	for _, bm := range baseMenus {
		btns := make([]types.SysMenuBtn, 0, len(bm.MenuBtn))
		for _, btn := range bm.MenuBtn {
			btns = append(btns, types.SysMenuBtn{
				ID:       btn.ID,
				Name:     btn.Name,
				Desc:     btn.Desc,
				ButtonId: btn.ID,
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
			Hidden: bm.Hidden,
			Button: btns,
		})
	}

	return &types.MenuTreeRes{
		MenuTree: buildMenuTree(menus),
	}, nil
}
