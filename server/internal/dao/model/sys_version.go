package model

// SysVersion 版本管理
type SysVersion struct {
	GvaModel
	VersionName *string `json:"versionName" gorm:"size:255;comment:版本名称"`
	VersionCode *string `json:"versionCode" gorm:"size:100;comment:版本号"`
	Description *string `json:"description" gorm:"size:500;comment:版本描述"`
	VersionData *string `json:"versionData" gorm:"type:text;comment:版本数据JSON"`
}

func (SysVersion) TableName() string {
	return "sys_versions"
}
