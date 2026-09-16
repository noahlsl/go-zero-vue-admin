// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package attachment

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FindAttachmentCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAttachmentCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAttachmentCategoryLogic {
	return &FindAttachmentCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAttachmentCategoryLogic) FindAttachmentCategory(req *types.GetById) (resp *types.ExaAttachmentCategoryRes, err error) {
	var category model.ExaAttachmentCategory
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("附件分类不存在")
		}
		return nil, errors.Wrap(err, "查询附件分类详情失败")
	}
	resp = &types.ExaAttachmentCategoryRes{
		Category: types.ExaAttachmentCategory{
			ID:   category.ID,
			Name: category.Name,
			Pid:  category.Pid,
		},
	}
	return resp, nil
}
