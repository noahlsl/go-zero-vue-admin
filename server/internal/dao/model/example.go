package model

// ExaCustomer 客户
type ExaCustomer struct {
	GvaModel
	CustomerName       string `json:"customerName" gorm:"comment:客户名"`
	CustomerPhoneData  string `json:"customerPhoneData" gorm:"comment:客户手机号"`
	SysUserID          uint   `json:"sysUserId" gorm:"comment:管理ID"`
	SysUserAuthorityID uint   `json:"sysUserAuthorityId" gorm:"comment:管理角色ID"`
	// -:migration 跳过该关联的外键迁移，避免 AutoMigrate 生成额外约束；关联本身仍可用于 Preload
	SysUser SysUser `json:"sysUser" gorm:"foreignKey:SysUserID;-:migration"`
}

func (ExaCustomer) TableName() string {
	return "exa_customers"
}

// ExaFileUploadAndDownload 文件上传下载
type ExaFileUploadAndDownload struct {
	GvaModel
	Name    string `json:"name" gorm:"comment:文件名"`
	ClassId int    `json:"classId" gorm:"default:0;comment:分类id"`
	Url     string `json:"url" gorm:"comment:文件地址"`
	Tag     string `json:"tag" gorm:"comment:文件标签"`
	Key     string `json:"key" gorm:"comment:编号"`
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}

// ExaAttachmentCategory 附件分类
type ExaAttachmentCategory struct {
	GvaModel
	Name string `json:"name" gorm:"type:varchar(255);comment:分类名称"`
	Pid  uint   `json:"pid" gorm:"default:0;comment:父节点ID"`
}

func (ExaAttachmentCategory) TableName() string {
	return "exa_attachment_category"
}

// ExaFile 断点续传文件
type ExaFile struct {
	GvaModel
	FileName   string `json:"fileName" gorm:"comment:文件名"`
	FileMd5    string `json:"fileMd5" gorm:"comment:文件MD5"`
	FilePath   string `json:"filePath" gorm:"comment:文件路径"`
	ChunkTotal int    `json:"chunkTotal" gorm:"comment:切片总数"`
	IsFinish   bool   `json:"isFinish" gorm:"comment:是否上传完成"`
}

func (ExaFile) TableName() string {
	return "exa_files"
}

// ExaFileChunk 断点续传文件切片
type ExaFileChunk struct {
	ID              uint   `gorm:"primarykey"`
	ExaFileID       uint   `json:"exaFileId" gorm:"comment:关联文件ID"`
	FileChunkNumber int    `json:"fileChunkNumber" gorm:"comment:切片编号"`
	FileChunkPath   string `json:"fileChunkPath" gorm:"comment:切片路径"`
}

func (ExaFileChunk) TableName() string {
	return "exa_file_chunks"
}
