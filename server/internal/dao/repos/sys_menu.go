package repos

import (
	"context"

	"zero/internal/dao/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysMenuRepos struct {
	*BaseRepos
}

func NewSysMenuRepos(db *gorm.DB) *SysMenuRepos {
	return &SysMenuRepos{BaseRepos: NewBaseRepos(db)}
}

func (r *SysMenuRepos) FindById(ctx context.Context, id uint) (*model.SysBaseMenu, error) {
	var menu model.SysBaseMenu
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&menu).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询菜单失败")
	}
	return &menu, nil
}
