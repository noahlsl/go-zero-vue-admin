package model

// SysDictionary 字典
type SysDictionary struct {
	GvaModel
	Name     string          `json:"name" gorm:"comment:字典名（中）"`
	Type     string          `json:"type" gorm:"comment:字典名（英）"`
	Status   *bool           `json:"status" gorm:"default:true;comment:状态"`
	Desc     string          `json:"desc" gorm:"comment:描述"`
	ParentID *uint           `json:"parentID" gorm:"comment:父级字典ID"`
	Children []SysDictionary `json:"children" gorm:"foreignKey:ParentID"`
}

func (SysDictionary) TableName() string {
	return "sys_dictionaries"
}

// SysDictionaryDetail 字典详情
type SysDictionaryDetail struct {
	GvaModel
	Label           string                `json:"label" gorm:"comment:展示值"`
	Value           string                `json:"value" gorm:"comment:字典值"`
	Extend          string                `json:"extend" gorm:"comment:扩展值"`
	Status          *bool                 `json:"status" gorm:"default:true;comment:启用状态"`
	Sort            int                   `json:"sort" gorm:"comment:排序标记"`
	SysDictionaryID int                   `json:"sysDictionaryID" gorm:"comment:关联标记"`
	ParentID        *uint                 `json:"parentID" gorm:"comment:父级字典详情ID"`
	Children        []SysDictionaryDetail `json:"children" gorm:"foreignKey:ParentID"`
	Level           int                   `json:"level" gorm:"comment:层级深度"`
	Path            string                `json:"path" gorm:"comment:层级路径"`
	// Disabled 非数据库字段，由 status 取反计算得出，供前端树形控件禁用节点
	Disabled bool `json:"disabled" gorm:"-"`
}

func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
