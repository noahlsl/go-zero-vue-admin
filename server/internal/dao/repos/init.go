package repos

import (
	"time"

	"zero/internal/dao/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitMySQL 初始化 MySQL 连接并自动迁移表结构
func InitMySQL(dsn string, maxIdleConns, maxOpenConns, connMaxLifetime int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	// AutoMigrate 所有表
	err = db.AutoMigrate(
		&model.SysUser{},
		&model.SysAuthority{},
		&model.SysBaseMenu{},
		&model.SysBaseMenuParameter{},
		&model.SysBaseMenuBtn{},
		&model.SysAuthorityBtn{},
		&model.SysApi{},
		&model.SysIgnoreApi{},
		&model.SysApiToken{},
		&model.SysAutoCodeHistory{},
		&model.SysAutoCodePackage{},
		&model.SysAIWorkflowSession{},
		&model.SysDictionary{},
		&model.SysDictionaryDetail{},
		&model.SysError{},
		&model.SysExportTemplate{},
		&model.SysExportTemplateCondition{},
		&model.SysExportTemplateJoin{},
		&model.SysCasbinRule{},
		&model.SysJwtBlacklist{},
		&model.SysLoginLog{},
		&model.SysOperationRecord{},
		&model.SysParams{},
		&model.SysVersion{},
		&model.ExaCustomer{},
		&model.ExaFileUploadAndDownload{},
		&model.ExaAttachmentCategory{},
		&model.ExaFile{},
		&model.ExaFileChunk{},
		&model.SysSkill{},
		&model.SysSkillReference{},
		&model.SysSkillResource{},
		&model.SysSkillScript{},
		&model.SysSkillTemplate{},
		&model.SysSkillGlobalConstraint{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
