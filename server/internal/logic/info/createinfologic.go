package info

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInfoLogic {
	return &CreateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateInfo 创建公告信息
func (l *CreateInfoLogic) CreateInfo(req *types.CreateAnnouncementInfoReq) (resp *types.Response, err error) {
	announcement := model.AnnouncementsInfo{
		Title:       req.Title,
		Content:     req.Content,
		UserID:      req.UserID,
		Attachments: []byte(req.Attachments),
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&announcement).Error; err != nil {
		return nil, errors.Wrap(err, "创建公告失败")
	}

	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
