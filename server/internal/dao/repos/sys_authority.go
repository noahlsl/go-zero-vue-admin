package repos

import (
	"context"

	"zero/internal/dao/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysAuthorityRepos struct {
	*BaseRepos
}

func NewSysAuthorityRepos(db *gorm.DB) *SysAuthorityRepos {
	return &SysAuthorityRepos{BaseRepos: NewBaseRepos(db)}
}

func (r *SysAuthorityRepos) FindById(ctx context.Context, id uint) (*model.SysAuthority, error) {
	var auth model.SysAuthority
	err := r.DB.WithContext(ctx).Where("authority_id = ?", id).First(&auth).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询角色失败")
	}
	return &auth, nil
}
