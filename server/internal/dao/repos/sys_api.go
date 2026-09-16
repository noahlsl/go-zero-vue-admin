package repos

import (
	"context"

	"zero/internal/dao/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysApiRepos struct {
	*BaseRepos
}

func NewSysApiRepos(db *gorm.DB) *SysApiRepos {
	return &SysApiRepos{BaseRepos: NewBaseRepos(db)}
}

func (r *SysApiRepos) FindById(ctx context.Context, id uint) (*model.SysApi, error) {
	var api model.SysApi
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&api).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询API失败")
	}
	return &api, nil
}
