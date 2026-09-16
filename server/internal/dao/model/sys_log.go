package model

import "time"

// SysLoginLog 登录日志
type SysLoginLog struct {
	GvaModel
	Username     string `json:"username" gorm:"comment:用户名"`
	Ip           string `json:"ip" gorm:"comment:请求ip"`
	Status       bool   `json:"status" gorm:"comment:登录状态"`
	ErrorMessage string `json:"errorMessage" gorm:"comment:错误信息"`
	Agent        string `json:"agent" gorm:"comment:代理"`
	UserID       uint   `json:"userId" gorm:"comment:用户id"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_logs"
}

// SysOperationRecord 操作日志
type SysOperationRecord struct {
	GvaModel
	Ip           string        `json:"ip" gorm:"comment:请求ip"`
	Method       string        `json:"method" gorm:"comment:请求方法"`
	Path         string        `json:"path" gorm:"comment:请求路径"`
	Status       int           `json:"status" gorm:"comment:请求状态"`
	Latency      time.Duration `json:"latency" gorm:"comment:延迟"`
	Agent        string        `json:"agent" gorm:"type:text;comment:代理"`
	ErrorMessage string        `json:"errorMessage" gorm:"comment:错误信息"`
	Body         string        `json:"body" gorm:"type:text;comment:请求Body"`
	Resp         string        `json:"resp" gorm:"type:text;comment:响应Body"`
	UserID       int           `json:"userId" gorm:"comment:用户id"`
}

func (SysOperationRecord) TableName() string {
	return "sys_operation_records"
}

// SysError 错误日志
type SysError struct {
	GvaModel
	Form     *string `json:"form" gorm:"type:text;comment:错误来源"`
	Info     *string `json:"info" gorm:"type:text;comment:错误内容"`
	Level    string  `json:"level" gorm:"comment:日志等级"`
	Solution *string `json:"solution" gorm:"type:text;comment:解决方案"`
	Status   string  `json:"status" gorm:"default:未处理;comment:处理状态"`
}

func (SysError) TableName() string {
	return "sys_error"
}
