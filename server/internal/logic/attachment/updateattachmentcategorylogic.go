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

type UpdateAttachmentCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAttachmentCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttachmentCategoryLogic {
	return &UpdateAttachmentCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAttachmentCategoryLogic) UpdateAttachmentCategory(req *types.ExaAttachmentCategory) (resp *types.Response, err error) {
	// 检查同级下是否已存在同名分类（排除自身）
	var count int64
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.ExaAttachmentCategory{}).
		Where("name = ? AND pid = ? AND id != ?", req.Name, req.Pid, req.ID).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询分类名称失败")
	}
	if count > 0 {
		return nil, errors.New("分类名称已存在")
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.ExaAttachmentCategory{}).
		Where("id = ?", req.ID).
		Updates(&model.ExaAttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
		logx.Errorw("更新附件分类失败",
			logx.Field("error", err.Error()),
			logx.Field("id", req.ID),
			logx.Field("module", "attachment"),
			logx.Field("action", "updateAttachmentCategory"),
		)
		return nil, errors.Wrap(err, "更新附件分类失败")
	}
	return &types.Response{}, nil
}
