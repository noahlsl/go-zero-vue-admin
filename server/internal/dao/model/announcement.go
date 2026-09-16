package model

import "gorm.io/datatypes"

// AnnouncementsInfo 公告信息
type AnnouncementsInfo struct {
	GvaModel
	Title       string         `json:"title" gorm:"column:title;comment:公告标题"`
	Content     string         `json:"content" gorm:"column:content;type:text;comment:公告内容"`
	UserID      *int           `json:"userID" gorm:"column:user_id;comment:发布者"`
	Attachments datatypes.JSON `json:"attachments" gorm:"column:attachments;type:json;comment:相关附件"`
}

func (AnnouncementsInfo) TableName() string {
	return "gva_announcements_info"
}
