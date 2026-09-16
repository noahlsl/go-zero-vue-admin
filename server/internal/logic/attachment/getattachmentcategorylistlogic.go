// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package attachment

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttachmentCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAttachmentCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttachmentCategoryListLogic {
	return &GetAttachmentCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAttachmentCategoryList 获取全部附件分类，并按父子关系组装为树形结构。
// 原 Gin 接口无入参，前端不传请求参数，且直接以数组形式返回 data。
func (l *GetAttachmentCategoryListLogic) GetAttachmentCategoryList() (resp []types.ExaAttachmentCategory, err error) {
	var categories []model.ExaAttachmentCategory
	if err = l.svcCtx.DB.WithContext(l.ctx).Find(&categories).Error; err != nil {
		logx.Errorw("查询附件分类列表失败",
			logx.Field("error", err.Error()),
			logx.Field("module", "attachment"),
			logx.Field("action", "getAttachmentCategoryList"),
		)
		return nil, errors.Wrap(err, "查询附件分类列表失败")
	}

	return buildCategoryTree(categories, 0), nil
}

// buildCategoryTree 递归构建分类树
func buildCategoryTree(categories []model.ExaAttachmentCategory, parentID uint) []types.ExaAttachmentCategory {
	tree := make([]types.ExaAttachmentCategory, 0)
	for _, c := range categories {
		if c.Pid != parentID {
			continue
		}
		tree = append(tree, types.ExaAttachmentCategory{
			ID:        c.ID,
			CreatedAt: c.CreatedAt.Format(time.RFC3339),
			UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
			Name:      c.Name,
			Pid:       c.Pid,
			Children:  buildCategoryTree(categories, c.ID),
		})
	}
	return tree
}
