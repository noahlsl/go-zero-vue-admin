package conv

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// defaultRouterFallback 角色默认路由无效时的兜底路由，与原 Gin 后端保持一致
const defaultRouterFallback = "404"

// ToTypesUser 将 model.SysUser 转换为 types.SysUser。
// 覆盖前端登录态依赖的全部字段（uuid、当前角色、角色列表、界面配置），
// 使接口返回结构与原 Gin 后端直接序列化 model 的结果一致。
func ToTypesUser(u model.SysUser) types.SysUser {
	return types.SysUser{
		ID:            u.ID,
		CreatedAt:     u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     u.UpdatedAt.Format(time.RFC3339),
		UUID:          u.UUID.String(),
		Username:      u.Username,
		NickName:      u.NickName,
		HeaderImg:     u.HeaderImg,
		AuthorityId:   u.AuthorityId,
		Authority:     ToTypesAuthority(u.Authority),
		Authorities:   ToTypesAuthorities(u.Authorities),
		Phone:         u.Phone,
		Email:         u.Email,
		Enable:        u.Enable,
		OriginSetting: u.OriginSetting,
	}
}

// ToTypesAuthority 将 model.SysAuthority 转换为 types.SysAuthority
func ToTypesAuthority(a model.SysAuthority) types.SysAuthority {
	ta := types.SysAuthority{
		CreatedAt:     a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     a.UpdatedAt.Format(time.RFC3339),
		AuthorityId:   a.AuthorityId,
		AuthorityName: a.AuthorityName,
		DefaultRouter: a.DefaultRouter,
		DataScope:     a.DataScope,
	}
	if a.ParentId != nil {
		ta.ParentId = *a.ParentId
	}
	return ta
}

// ToTypesAuthorities 批量转换角色列表
func ToTypesAuthorities(as []model.SysAuthority) []types.SysAuthority {
	if len(as) == 0 {
		return nil
	}

	list := make([]types.SysAuthority, 0, len(as))
	for _, a := range as {
		list = append(list, ToTypesAuthority(a))
	}
	return list
}

// EnsureDefaultRouter 校验角色默认路由是否在角色已授权的菜单范围内，
// 不在范围内时回退为 404，与原 Gin 后端 UserAuthorityDefaultRouter 行为一致。
// 校验失败不阻断主流程，仅记录日志并保持兜底路由。
func EnsureDefaultRouter(ctx context.Context, db *gorm.DB, authority *model.SysAuthority) {
	if authority == nil {
		return
	}

	var menuIds []string
	if err := db.WithContext(ctx).Model(&model.SysAuthorityMenu{}).
		Where("sys_authority_authority_id = ?", authority.AuthorityId).
		Pluck("sys_base_menu_id", &menuIds).Error; err != nil {
		logx.WithContext(ctx).Errorw("查询角色菜单失败",
			logx.Field("error", err.Error()),
			logx.Field("authority_id", authority.AuthorityId),
			logx.Field("module", "user"),
			logx.Field("action", "ensure_default_router"),
		)
		return
	}

	var count int64
	if err := db.WithContext(ctx).Model(&model.SysBaseMenu{}).
		Where("name = ? AND id IN (?)", authority.DefaultRouter, menuIds).
		Count(&count).Error; err != nil {
		logx.WithContext(ctx).Errorw("校验角色默认路由失败",
			logx.Field("error", err.Error()),
			logx.Field("authority_id", authority.AuthorityId),
			logx.Field("module", "user"),
			logx.Field("action", "ensure_default_router"),
		)
		return
	}

	if count == 0 {
		authority.DefaultRouter = defaultRouterFallback
	}
}
