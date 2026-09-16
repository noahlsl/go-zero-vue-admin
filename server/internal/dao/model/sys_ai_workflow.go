package model

// SysAIWorkflowSession AI 工作流会话
type SysAIWorkflowSession struct {
	GvaModel
	UserID         uint           `json:"userId" gorm:"comment:用户ID"`
	Tab            string         `json:"tab" gorm:"size:32;comment:会话类型"`
	Title          string         `json:"title" gorm:"size:255;comment:会话标题"`
	Summary        string         `json:"summary" gorm:"type:text;comment:摘要"`
	ConversationID string         `json:"conversationId" gorm:"comment:Dify会话ID"`
	MessageID      string         `json:"messageId" gorm:"comment:Dify消息ID"`
	CurrentNodeID  string         `json:"currentNodeId" gorm:"size:64;comment:当前选中节点ID"`
	Settings       map[string]any `json:"settings" gorm:"type:text;serializer:json;comment:页面设置"`
	FormData       map[string]any `json:"formData" gorm:"type:text;serializer:json;comment:表单数据"`
	ResultData     map[string]any `json:"resultData" gorm:"type:text;serializer:json;comment:当前展示结果"`
	Messages       map[string]any `json:"messages" gorm:"type:text;serializer:json;comment:会话消息"`
}

func (SysAIWorkflowSession) TableName() string {
	return "sys_ai_workflow_sessions"
}
