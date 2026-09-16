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
)

type DeleteAttachmentCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAttachmentCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAttachmentCategoryLogic {
	return &DeleteAttachmentCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAttachmentCategoryLogic) DeleteAttachmentCategory(req *types.GetById) (resp *types.Response, err error) {
	// 检查是否存在子分类
	var childCount int64
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.ExaAttachmentCategory{}).
		Where("pid = ?", req.ID).
		Count(&childCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询子分类失败")
	}
	if childCount > 0 {
		return nil, errors.New("请先删除子级")
	}

	// 使用 Unscoped 进行硬删除，与原 Gin 逻辑保持一致
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Unscoped().
		Where("id = ?", req.ID).
		Delete(&model.ExaAttachmentCategory{}).Error; err != nil {
		logx.Errorw("删除附件分类失败",
			logx.Field("error", err.Error()),
			logx.Field("id", req.ID),
			logx.Field("module", "attachment"),
			logx.Field("action", "deleteAttachmentCategory"),
		)
		return nil, errors.Wrap(err, "删除附件分类失败")
	}
	return &types.Response{}, nil
}
