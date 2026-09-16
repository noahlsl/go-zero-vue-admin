package model

// SysCasbinRule Casbin 策略规则
type SysCasbinRule struct {
	Ptype string `json:"ptype" gorm:"type:varchar(100);comment:策略类型"`
	V0    string `json:"v0" gorm:"type:varchar(100);comment:角色ID"`
	V1    string `json:"v1" gorm:"type:varchar(255);comment:请求路径"`
	V2    string `json:"v2" gorm:"type:varchar(100);comment:请求方法"`
	V3    string `json:"v3" gorm:"type:varchar(100);comment:保留字段"`
	V4    string `json:"v4" gorm:"type:varchar(100);comment:保留字段"`
	V5    string `json:"v5" gorm:"type:varchar(100);comment:保留字段"`
}

func (SysCasbinRule) TableName() string {
	return "sys_casbin_rule"
}
