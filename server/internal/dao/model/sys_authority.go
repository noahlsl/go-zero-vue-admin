package model

import (
	"time"

	"gorm.io/gorm"
)

type SysAuthority struct {
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	AuthorityId   uint           `json:"authorityId" gorm:"not null;unique;primary_key;size:90"`
	AuthorityName string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId      *uint          `json:"parentId" gorm:"comment:父角色ID"`
	DefaultRouter string         `json:"defaultRouter" gorm:"comment:默认菜单"`
	DataScope     string         `json:"dataScope" gorm:"comment:数据范围"`
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
