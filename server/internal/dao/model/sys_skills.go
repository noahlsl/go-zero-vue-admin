package model

// SysSkill 技能管理
type SysSkill struct {
	GvaModel
	Name    string `json:"name" gorm:"size:255;comment:技能名称"`
	Desc    string `json:"desc" gorm:"size:500;comment:技能描述"`
	Content string `json:"content" gorm:"type:text;comment:技能内容"`
}

func (SysSkill) TableName() string {
	return "sys_skills"
}

// SysSkillReference 技能参考资料
type SysSkillReference struct {
	GvaModel
	Name    string `json:"name" gorm:"size:255;comment:文件名"`
	Content string `json:"content" gorm:"type:text;comment:内容"`
}

func (SysSkillReference) TableName() string {
	return "sys_skill_references"
}

// SysSkillResource 技能资源文件
type SysSkillResource struct {
	GvaModel
	Name    string `json:"name" gorm:"size:255;comment:文件名"`
	Type    string `json:"type" gorm:"size:50;comment:资源类型"`
	Content string `json:"content" gorm:"type:text;comment:内容"`
}

func (SysSkillResource) TableName() string {
	return "sys_skill_resources"
}

// SysSkillScript 技能脚本
type SysSkillScript struct {
	GvaModel
	Name    string `json:"name" gorm:"size:255;comment:脚本名"`
	Content string `json:"content" gorm:"type:text;comment:脚本内容"`
}

func (SysSkillScript) TableName() string {
	return "sys_skill_scripts"
}

// SysSkillTemplate 技能模板
type SysSkillTemplate struct {
	GvaModel
	Name    string `json:"name" gorm:"size:255;comment:模板名"`
	Content string `json:"content" gorm:"type:text;comment:模板内容"`
}

func (SysSkillTemplate) TableName() string {
	return "sys_skill_templates"
}

// SysSkillGlobalConstraint 全局约束
type SysSkillGlobalConstraint struct {
	GvaModel
	Content string `json:"content" gorm:"type:text;comment:约束内容"`
}

func (SysSkillGlobalConstraint) TableName() string {
	return "sys_skill_global_constraints"
}
