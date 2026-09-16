package model

// SysApi API 管理
type SysApi struct {
	GvaModel
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"api_group" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:POST;comment:方法"`
}

func (SysApi) TableName() string {
	return "sys_apis"
}

// SysIgnoreApi 忽略的 API
type SysIgnoreApi struct {
	GvaModel
	Path   string `json:"path" gorm:"comment:api路径"`
	Method string `json:"method" gorm:"default:POST;comment:方法"`
}

func (SysIgnoreApi) TableName() string {
	return "sys_ignore_apis"
}
