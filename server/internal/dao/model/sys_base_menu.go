package model

import (
	"time"

	"gorm.io/gorm"
)

// SysBaseMenu 基础菜单
type SysBaseMenu struct {
	GvaModel
	MenuLevel      uint                   `json:"-" gorm:"comment:菜单层级"`
	ParentId       uint                   `json:"parentId" gorm:"comment:父菜单ID"`
	Path           string                 `json:"path" gorm:"comment:路由path"`
	Name           string                 `json:"name" gorm:"comment:路由name"`
	Hidden         bool                   `json:"hidden" gorm:"default:false;comment:是否在列表隐藏"`
	Component      string                 `json:"component" gorm:"comment:对应前端文件路径"`
	Sort           int                    `json:"sort" gorm:"comment:排序标记"`
	ActiveName     string                 `json:"activeName" gorm:"column:meta_active_name"`
	KeepAlive      bool                   `json:"keepAlive" gorm:"column:meta_keep_alive"`
	DefaultMenu    bool                   `json:"defaultMenu" gorm:"column:meta_default_menu"`
	Title          string                 `json:"title" gorm:"column:meta_title"`
	Icon           string                 `json:"icon" gorm:"column:meta_icon"`
	CloseTab       bool                   `json:"closeTab" gorm:"column:meta_close_tab"`
	TransitionType string                 `json:"transitionType" gorm:"column:meta_transition_type"`
	Parameters     []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID"`
	MenuBtn        []SysBaseMenuBtn       `json:"menuBtn" gorm:"foreignKey:SysBaseMenuID"`
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}

// SysBaseMenuParameter 菜单参数
type SysBaseMenuParameter struct {
	ID            uint   `gorm:"primarykey"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:关联菜单ID"`
	Type          string `json:"type" gorm:"comment:params or query"`
	Key           string `json:"key" gorm:"comment:参数key"`
	Value         string `json:"value" gorm:"comment:参数值"`
}

func (SysBaseMenuParameter) TableName() string {
	return "sys_base_menu_parameters"
}

// SysBaseMenuBtn 菜单按钮
type SysBaseMenuBtn struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"comment:按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}

func (SysBaseMenuBtn) TableName() string {
	return "sys_base_menu_btns"
}

// SysAuthorityBtn 角色按钮权限
type SysAuthorityBtn struct {
	ID               uint           `gorm:"primarykey"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	AuthorityId      uint           `json:"authorityId" gorm:"comment:角色ID"`
	SysMenuID        uint           `json:"sysMenuId" gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `json:"sysBaseMenuBtnId" gorm:"comment:菜单按钮ID"`
}

func (SysAuthorityBtn) TableName() string {
	return "sys_authority_btns"
}
