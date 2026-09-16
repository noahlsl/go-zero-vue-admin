package model

// SysAutoCodeHistory 代码生成历史
type SysAutoCodeHistory struct {
	GvaModel
	Table            string            `json:"table" gorm:"comment:表名"`
	Package          string            `json:"package" gorm:"comment:模块名"`
	Request          string            `json:"request" gorm:"type:text;comment:前端传入的结构化信息"`
	StructName       string            `json:"structName" gorm:"comment:结构体名称"`
	Abbreviation     string            `json:"abbreviation" gorm:"comment:结构体名称缩写"`
	BusinessDB       string            `json:"businessDB" gorm:"comment:业务库"`
	Description      string            `json:"description" gorm:"comment:Struct中文名称"`
	Templates        map[string]string `json:"templates" gorm:"serializer:json;comment:模板信息"`
	Injections       map[string]string `json:"injections" gorm:"serializer:json;comment:注入路径"`
	Flag             int               `json:"flag" gorm:"comment:0创建 1回滚"`
	ApiIDs           []uint            `json:"apiIDs" gorm:"serializer:json;comment:api表注册内容"`
	MenuID           uint              `json:"menuID" gorm:"comment:菜单ID"`
	ExportTemplateID uint              `json:"exportTemplateID" gorm:"comment:导出模板ID"`
	PackageID        uint              `json:"packageID" gorm:"comment:包ID"`
}

func (SysAutoCodeHistory) TableName() string {
	return "sys_auto_code_histories"
}

// SysAutoCodePackage 代码生成包
type SysAutoCodePackage struct {
	GvaModel
	Desc        string `json:"desc" gorm:"comment:描述"`
	Label       string `json:"label" gorm:"comment:展示名"`
	Template    string `json:"template" gorm:"comment:模版"`
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Module      string `json:"-" gorm:"comment:模块"`
}

func (SysAutoCodePackage) TableName() string {
	return "sys_auto_code_packages"
}
