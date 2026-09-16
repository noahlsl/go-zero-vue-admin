package info

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInfoLogic {
	return &DeleteInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteInfo 删除公告信息
func (l *DeleteInfoLogic) DeleteInfo(req *types.DeleteAnnouncementInfoReq) (resp *types.Response, err error) {
	var announcement model.AnnouncementsInfo
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&announcement).Error; err != nil {
		return nil, errors.Wrap(err, "查询公告失败")
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Delete(&announcement).Error; err != nil {
		return nil, errors.Wrap(err, "删除公告失败")
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}
