package menu

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRolesLogic {
	return &GetMenuRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetMenuRoles 获取拥有指定菜单的角色ID列表
// 对齐原 Gin 实现：authorityIds 来自 sys_authority_menus 关联表，
// defaultRouterAuthorityIds 来自将菜单名设为默认首页（default_router = menu.Name）的角色
func (l *GetMenuRolesLogic) GetMenuRoles(req *types.SetMenuRolesReq) (resp *types.MenuRolesRes, err error) {
	if req.MenuId == 0 {
		return nil, errors.New("菜单ID不能为空")
	}

	authorityIds, err := l.getAuthorityIdsByMenuId(req.MenuId)
	if err != nil {
		return nil, err
	}

	defaultRouterAuthorityIds, err := l.getDefaultRouterAuthorityIds(req.MenuId)
	if err != nil {
		return nil, err
	}

	return &types.MenuRolesRes{
		AuthorityIds:              authorityIds,
		DefaultRouterAuthorityIds: defaultRouterAuthorityIds,
	}, nil
}

// getAuthorityIdsByMenuId 查询拥有指定菜单的角色ID列表
func (l *GetMenuRolesLogic) getAuthorityIdsByMenuId(menuId uint) ([]uint, error) {
	var records []model.SysAuthorityMenu
	if err := l.svcCtx.DB.WithContext(l.ctx).Where("sys_base_menu_id = ?", menuId).
		Find(&records).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单角色关联失败")
	}

	authorityIds := make([]uint, 0, len(records))
	for _, r := range records {
		id, e := strconv.Atoi(r.AuthorityId)
		if e == nil {
			authorityIds = append(authorityIds, uint(id))
		}
	}

	return authorityIds, nil
}

// getDefaultRouterAuthorityIds 查询将指定菜单设为首页的角色ID列表
func (l *GetMenuRolesLogic) getDefaultRouterAuthorityIds(menuId uint) ([]uint, error) {
	var menu model.SysBaseMenu
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&menu, menuId).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单信息失败")
	}

	var authorities []model.SysAuthority
	if err := l.svcCtx.DB.WithContext(l.ctx).Where("default_router = ?", menu.Name).
		Find(&authorities).Error; err != nil {
		return nil, errors.Wrap(err, "查询默认首页角色失败")
	}

	authorityIds := make([]uint, 0, len(authorities))
	for _, auth := range authorities {
		authorityIds = append(authorityIds, auth.AuthorityId)
	}

	return authorityIds, nil
}
