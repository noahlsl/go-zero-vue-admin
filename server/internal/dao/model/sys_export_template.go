package model

// SysExportTemplate 导出模板
type SysExportTemplate struct {
	GvaModel
	DBName       string `json:"dbName" gorm:"comment:数据库名称"`
	Name         string `json:"name" gorm:"comment:模板名称"`
	Table        string `json:"tableName" gorm:"column:table_name;comment:表名称"`
	TemplateID   string `json:"templateID" gorm:"comment:模板标识"`
	TemplateInfo string `json:"templateInfo" gorm:"type:text;comment:模板信息"`
	SQL          string `json:"sql" gorm:"type:text;comment:自定义导出SQL"`
	ImportSQL    string `json:"importSQL" gorm:"type:text;comment:自定义导入SQL"`
	Limit        *int   `json:"limit" gorm:"comment:导出限制"`
	Order        string `json:"order" gorm:"comment:排序"`
}

func (SysExportTemplate) TableName() string {
	return "sys_export_templates"
}

// SysExportTemplateCondition 导出模板条件
type SysExportTemplateCondition struct {
	ID         uint   `gorm:"primarykey"`
	TemplateID string `json:"templateID" gorm:"comment:模板标识"`
	From       string `json:"from" gorm:"comment:来源"`
	Column     string `json:"column" gorm:"comment:列名"`
	Operator   string `json:"operator" gorm:"comment:操作符"`
}

func (SysExportTemplateCondition) TableName() string {
	return "sys_export_template_conditions"
}

// SysExportTemplateJoin 导出模板关联
type SysExportTemplateJoin struct {
	ID         uint   `gorm:"primarykey"`
	TemplateID string `json:"templateID" gorm:"comment:模板标识"`
	JOINS      string `json:"joins" gorm:"comment:关联条件"`
	Table      string `json:"table" gorm:"comment:关联表"`
	ON         string `json:"on" gorm:"comment:关联字段"`
}

func (SysExportTemplateJoin) TableName() string {
	return "sys_export_template_joins"
}
