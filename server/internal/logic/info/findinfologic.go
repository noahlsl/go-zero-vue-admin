package info

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindInfoLogic {
	return &FindInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FindInfo 根据ID查找公告信息
func (l *FindInfoLogic) FindInfo(req *types.FindAnnouncementInfoReq) (resp *types.AnnouncementInfoRes, err error) {
	if req.ID == nil {
		return nil, errors.New("公告ID不能为空")
	}

	var announcement model.AnnouncementsInfo
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", *req.ID).First(&announcement).Error; err != nil {
		return nil, errors.Wrap(err, "查询公告失败")
	}

	return &types.AnnouncementInfoRes{
		Info: conv.ToTypesAnnouncementInfo(announcement),
	}, nil
}
