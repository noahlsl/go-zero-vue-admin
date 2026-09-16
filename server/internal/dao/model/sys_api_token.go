package model

import "time"

type SysApiToken struct {
	GvaModel
	UserID      uint      `json:"userId" gorm:"comment:用户ID"`
	AuthorityId uint      `json:"authorityId" gorm:"comment:角色ID"`
	Token       string    `json:"token" gorm:"type:text;comment:Token"`
	Status      bool      `json:"status" gorm:"comment:状态，true有效"`
	ExpiresAt   time.Time `json:"expiresAt" gorm:"comment:过期时间"`
	Remark      string    `json:"remark" gorm:"comment:备注"`
}

func (SysApiToken) TableName() string {
	return "sys_api_tokens"
}
