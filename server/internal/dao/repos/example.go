package repos

import (
	"context"

	"zero/internal/dao/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ExaCustomerRepos struct {
	*BaseRepos
}

func NewExaCustomerRepos(db *gorm.DB) *ExaCustomerRepos {
	return &ExaCustomerRepos{BaseRepos: NewBaseRepos(db)}
}

func (r *ExaCustomerRepos) FindById(ctx context.Context, id uint) (*model.ExaCustomer, error) {
	var customer model.ExaCustomer
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询客户失败")
	}
	return &customer, nil
}
