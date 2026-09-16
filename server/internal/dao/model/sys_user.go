package model

import "github.com/google/uuid"

type SysUser struct {
	GvaModel
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`
	Password    string    `json:"-" gorm:"comment:用户登录密码"`
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	AuthorityId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	// -:migration 跳过该关联的外键迁移：GORM 会因两侧同名字段 AuthorityId 误判关联方向，
	// 生成 sys_authorities -> sys_users 的反向外键导致 AutoMigrate 失败；关联本身仍可用于 Preload
	Authority     SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色;-:migration"`
	Authorities   []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone         string         `json:"phone" gorm:"comment:用户手机号"`
	Email         string         `json:"email" gorm:"comment:用户邮箱"`
	Enable        int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	OriginSetting map[string]any `json:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
