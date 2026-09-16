package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserAuthorityLogic {
	return &SetUserAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserAuthorityLogic) SetUserAuthority(req *types.SetUserAuthReq) (resp *types.Response, err error) {
	// 1. 从 context 获取当前登录用户 ID
	userId := middleware.GetUserId(l.ctx)
	if userId == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	// 2. 校验用户是否拥有该角色
	var userAuth model.SysUserAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("sys_user_id = ? AND sys_authority_authority_id = ?",
		userId, req.AuthorityId).First(&userAuth).Error; err != nil {
		return nil, errors.New("该用户无此角色")
	}

	// 3. 查询角色信息
	var authority model.SysAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", req.AuthorityId).First(&authority).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色信息失败")
	}

	// 4. 查询角色关联的菜单，校验默认路由是否存在
	var authorityMenus []model.SysAuthorityMenu
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("sys_authority_authority_id = ?", req.AuthorityId).
		Find(&authorityMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色菜单失败")
	}

	menuIds := make([]string, 0, len(authorityMenus))
	for _, am := range authorityMenus {
		menuIds = append(menuIds, am.MenuId)
	}

	var baseMenus []model.SysBaseMenu
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id in (?)", menuIds).
		Find(&baseMenus).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单详情失败")
	}

	hasMenu := false
	for _, m := range baseMenus {
		if m.Name == authority.DefaultRouter {
			hasMenu = true
			break
		}
	}
	if !hasMenu {
		return nil, errors.New("找不到默认路由,无法切换本角色")
	}

	// 5. 更新用户的默认角色
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).Where("id = ?", userId).
		Update("authority_id", req.AuthorityId).Error; err != nil {
		l.Logger.Errorw("设置用户角色失败",
			logx.Field("error", err),
			logx.Field("user_id", userId),
			logx.Field("authority_id", req.AuthorityId),
			logx.Field("module", "user"),
			logx.Field("action", "set_user_authority"),
		)
		return nil, errors.Wrap(err, "设置用户角色失败")
	}

	return &types.Response{Msg: "修改成功"}, nil
}
