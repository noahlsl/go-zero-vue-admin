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

type DeleteAttachmentCategoryByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAttachmentCategoryByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAttachmentCategoryByIdsLogic {
	return &DeleteAttachmentCategoryByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAttachmentCategoryByIdsLogic) DeleteAttachmentCategoryByIds(req *types.IdsReq) (resp *types.Response, err error) {
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Unscoped().
		Where("id IN ?", req.Ids).
		Delete(&model.ExaAttachmentCategory{}).Error; err != nil {
		logx.Errorw("批量删除附件分类失败",
			logx.Field("error", err.Error()),
			logx.Field("ids", req.Ids),
			logx.Field("module", "attachment"),
			logx.Field("action", "deleteAttachmentCategoryByIds"),
		)
		return nil, errors.Wrap(err, "批量删除附件分类失败")
	}
	return &types.Response{}, nil
}
