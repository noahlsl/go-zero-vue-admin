package base

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	defaultAdminUsername = "admin"
	defaultAdminPassword = "123456"
	defaultAuthorityId   = 888
	defaultAuthorityName = "超级管理员"
	defaultRouter        = "dashboard"
)

type InitDBLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDBLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDBLogic {
	return &InitDBLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// InitDB 初始化数据库：创建默认角色和管理员账号
func (l *InitDBLogic) InitDB() error {
	db := l.svcCtx.DB.WithContext(l.ctx)

	// 检查是否已有管理员用户
	var count int64
	if err := db.Model(&model.SysUser{}).Count(&count).Error; err != nil {
		return errors.Wrap(err, "检查用户数据失败")
	}
	if count > 0 {
		return errors.New("数据库已初始化，请勿重复执行")
	}

	// 1. 创建默认角色
	if err := l.createDefaultAuthority(db); err != nil {
		return err
	}

	// 2. 创建管理员用户
	if err := l.createAdminUser(db); err != nil {
		return err
	}

	logx.Infow("数据库初始化成功", logx.Field("module", "initdb"))
	return nil
}

// createDefaultAuthority 创建默认超级管理员角色
func (l *InitDBLogic) createDefaultAuthority(db *gorm.DB) error {
	authority := model.SysAuthority{
		AuthorityId:   defaultAuthorityId,
		AuthorityName: defaultAuthorityName,
		DefaultRouter: defaultRouter,
	}
	err := db.Where("authority_id = ?", defaultAuthorityId).
		FirstOrCreate(&authority).Error
	if err != nil {
		return errors.Wrap(err, "创建默认角色失败")
	}
	return nil
}

// createAdminUser 创建默认管理员用户
func (l *InitDBLogic) createAdminUser(db *gorm.DB) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(defaultAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "加密密码失败")
	}

	user := model.SysUser{
		UUID:        uuid.New(),
		Username:    defaultAdminUsername,
		Password:    string(hashedPwd),
		NickName:    "超级管理员",
		HeaderImg:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
		AuthorityId: defaultAuthorityId,
		Enable:      1,
	}
	err = db.Create(&user).Error
	if err != nil {
		return errors.Wrap(err, "创建管理员用户失败")
	}
	return nil
}
