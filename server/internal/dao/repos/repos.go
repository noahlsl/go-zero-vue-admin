package repos

import "gorm.io/gorm"

// BaseRepos 通用数据库操作基类
type BaseRepos struct {
	DB *gorm.DB
}

// NewBaseRepos 创建基础 repo
func NewBaseRepos(db *gorm.DB) *BaseRepos {
	return &BaseRepos{DB: db}
}
