package sysDictionaryDetail

import "zero/internal/dao/model"

// setDisabled 按 status 计算 disabled 计算字段：
// status 为 false 时节点禁用；status 为空时默认不禁用。与原 Gin 后端行为保持一致。
func setDisabled(detail *model.SysDictionaryDetail) {
	if detail.Status != nil {
		detail.Disabled = !*detail.Status
		return
	}
	detail.Disabled = false
}
