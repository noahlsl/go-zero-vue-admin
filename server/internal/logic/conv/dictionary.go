package conv

import (
	"time"

	"zero/internal/dao/model"
	"zero/internal/types"
)

// ToTypesSysDictionary 将 model.SysDictionary 转换为 types.SysDictionary。
// 时间字段格式化为 RFC3339 字符串，使序列化结果与原 Gin 后端直接序列化 model 一致。
func ToTypesSysDictionary(d model.SysDictionary) types.SysDictionary {
	res := types.SysDictionary{
		ID:        d.ID,
		CreatedAt: d.CreatedAt.Format(time.RFC3339),
		UpdatedAt: d.UpdatedAt.Format(time.RFC3339),
		Name:      d.Name,
		Type:      d.Type,
		Status:    d.Status,
		Desc:      d.Desc,
		ParentID:  d.ParentID,
	}
	if d.Children != nil {
		res.Children = ToTypesSysDictionaries(d.Children)
	}
	return res
}

// ToTypesSysDictionaries 批量转换字典列表
func ToTypesSysDictionaries(ds []model.SysDictionary) []types.SysDictionary {
	if ds == nil {
		return nil
	}

	list := make([]types.SysDictionary, 0, len(ds))
	for _, d := range ds {
		list = append(list, ToTypesSysDictionary(d))
	}
	return list
}

// ToTypesSysDictionaryDetail 将 model.SysDictionaryDetail 转换为 types.SysDictionaryDetail。
// Disabled 为 gorm:"-" 计算字段，由调用方按 status 取反后写入，此处原样透传。
func ToTypesSysDictionaryDetail(d model.SysDictionaryDetail) types.SysDictionaryDetail {
	res := types.SysDictionaryDetail{
		ID:              d.ID,
		CreatedAt:       d.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       d.UpdatedAt.Format(time.RFC3339),
		Label:           d.Label,
		Value:           d.Value,
		Extend:          d.Extend,
		Status:          d.Status,
		Sort:            d.Sort,
		SysDictionaryID: d.SysDictionaryID,
		ParentID:        d.ParentID,
		Level:           d.Level,
		Path:            d.Path,
		Disabled:        d.Disabled,
	}
	if d.Children != nil {
		res.Children = ToTypesSysDictionaryDetails(d.Children)
	}
	return res
}

// ToTypesSysDictionaryDetails 批量转换字典详情列表
func ToTypesSysDictionaryDetails(ds []model.SysDictionaryDetail) []types.SysDictionaryDetail {
	if ds == nil {
		return nil
	}

	list := make([]types.SysDictionaryDetail, 0, len(ds))
	for _, d := range ds {
		list = append(list, ToTypesSysDictionaryDetail(d))
	}
	return list
}
