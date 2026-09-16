package conv

import (
	"time"

	"zero/internal/dao/model"
	"zero/internal/types"
)

// ToTypesAnnouncementInfo 将 model.AnnouncementsInfo 转换为 types.AnnouncementInfo。
// 时间字段格式化为 RFC3339 字符串，使序列化结果与原 Gin 后端直接序列化 model 一致。
func ToTypesAnnouncementInfo(a model.AnnouncementsInfo) types.AnnouncementInfo {
	return types.AnnouncementInfo{
		ID:          a.ID,
		CreatedAt:   a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   a.UpdatedAt.Format(time.RFC3339),
		Title:       a.Title,
		Content:     a.Content,
		UserID:      a.UserID,
		Attachments: string(a.Attachments),
	}
}

// ToTypesAnnouncementInfos 批量转换公告列表
func ToTypesAnnouncementInfos(as []model.AnnouncementsInfo) []types.AnnouncementInfo {
	if as == nil {
		return nil
	}

	list := make([]types.AnnouncementInfo, 0, len(as))
	for _, a := range as {
		list = append(list, ToTypesAnnouncementInfo(a))
	}
	return list
}
