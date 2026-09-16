package model

// SysJwtBlacklist JWT 黑名单
type SysJwtBlacklist struct {
	GvaModel
	Jwt string `json:"jwt" gorm:"type:text;comment:jwt"`
}

func (SysJwtBlacklist) TableName() string {
	return "jwt_blacklists"
}
