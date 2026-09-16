package model

import (
	"time"

	"gorm.io/gorm"
)

// GvaModel 基础模型，所有需要软删除的表都嵌入此结构
type GvaModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
