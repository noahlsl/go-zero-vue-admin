package repos

import (
	"context"

	"zero/internal/dao/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysUserRepos struct {
	*BaseRepos
}

func NewSysUserRepos(db *gorm.DB) *SysUserRepos {
	return &SysUserRepos{BaseRepos: NewBaseRepos(db)}
}

func (r *SysUserRepos) FindByUsername(ctx context.Context, username string) (*model.SysUser, error) {
	var user model.SysUser
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询用户失败")
	}
	return &user, nil
}

func (r *SysUserRepos) FindById(ctx context.Context, id uint) (*model.SysUser, error) {
	var user model.SysUser
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询用户失败")
	}
	return &user, nil
}
