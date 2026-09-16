package autoCode

import (
	"strings"
	"unicode"
)

// toFieldName 将下划线命名转换为驼峰命名（首字母大写）
// 例：sys_user -> SysUser, created_at -> CreatedAt
func toFieldName(name string) string {
	var result strings.Builder
	upperNext := true
	for _, r := range name {
		if r == '_' {
			upperNext = true
			continue
		}
		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// toFieldJson 将下划线命名保持原样（JSON 字段名）
// 例：sys_user -> sys_user
func toFieldJson(name string) string {
	return name
}
