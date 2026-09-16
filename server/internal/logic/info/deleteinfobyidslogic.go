package info

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInfoByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInfoByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInfoByIdsLogic {
	return &DeleteInfoByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteInfoByIds 批量删除公告信息
func (l *DeleteInfoByIdsLogic) DeleteInfoByIds(req *types.DeleteAnnouncementInfoByIdsReq) (resp *types.Response, err error) {
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("id IN ?", req.IDs).
		Delete(&model.AnnouncementsInfo{}).Error; err != nil {
		return nil, errors.Wrap(err, "批量删除公告失败")
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}
