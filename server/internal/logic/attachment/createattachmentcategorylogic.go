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

type CreateAttachmentCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAttachmentCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAttachmentCategoryLogic {
	return &CreateAttachmentCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAttachmentCategoryLogic) CreateAttachmentCategory(req *types.ExaAttachmentCategory) (resp *types.Response, err error) {
	// 检查同级下是否已存在同名分类
	var count int64
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.ExaAttachmentCategory{}).
		Where("name = ? AND pid = ?", req.Name, req.Pid).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询分类名称失败")
	}
	if count > 0 {
		return nil, errors.New("分类名称已存在")
	}

	category := model.ExaAttachmentCategory{
		Name: req.Name,
		Pid:  req.Pid,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&category).Error; err != nil {
		logx.Errorw("创建附件分类失败",
			logx.Field("error", err.Error()),
			logx.Field("module", "attachment"),
			logx.Field("action", "createAttachmentCategory"),
		)
		return nil, errors.Wrap(err, "创建附件分类失败")
	}
	return &types.Response{Msg: "创建/更新成功"}, nil
}
