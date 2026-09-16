package model

// SysParams 系统参数
type SysParams struct {
	GvaModel
	Name  string `json:"name" gorm:"comment:参数名称"`
	Key   string `json:"key" gorm:"comment:参数键"`
	Value string `json:"value" gorm:"comment:参数值"`
	Desc  string `json:"desc" gorm:"comment:参数说明"`
}

func (SysParams) TableName() string {
	return "sys_params"
}
