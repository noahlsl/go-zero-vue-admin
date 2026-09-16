package model

// SysUserAuthority 用户角色关联表
type SysUserAuthority struct {
	SysUserId               uint `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId uint `gorm:"column:sys_authority_authority_id"`
}

func (SysUserAuthority) TableName() string {
	return "sys_user_authority"
}

// SysAuthorityMenu 角色菜单关联表
type SysAuthorityMenu struct {
	MenuId      string `gorm:"column:sys_base_menu_id"`
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
}

func (SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
