package info

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateInfo 更新公告信息
func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateAnnouncementInfoReq) (resp *types.Response, err error) {
	var announcement model.AnnouncementsInfo
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&announcement).Error; err != nil {
		return nil, errors.Wrap(err, "查询公告失败")
	}

	updates := map[string]interface{}{
		"Title":       req.Title,
		"Content":     req.Content,
		"UserID":      req.UserID,
		"Attachments": []byte(req.Attachments),
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&announcement).Updates(updates).Error; err != nil {
		return nil, errors.Wrap(err, "更新公告失败")
	}

	return &types.Response{Code: 0, Msg: "修改成功"}, nil
}
