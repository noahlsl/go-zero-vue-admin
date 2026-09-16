package info

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoListLogic {
	return &GetInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetInfoList 分页查询公告列表
func (l *GetInfoListLogic) GetInfoList(req *types.GetAnnouncementInfoListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.AnnouncementsInfo{})

	// 时间范围过滤
	if req.StartCreatedAt != "" {
		startTime, parseErr := time.Parse(time.RFC3339, req.StartCreatedAt)
		if parseErr != nil {
			return nil, errors.Wrap(parseErr, "解析开始创建时间失败")
		}
		db = db.Where("created_at >= ?", startTime)
	}
	if req.EndCreatedAt != "" {
		endTime, parseErr := time.Parse(time.RFC3339, req.EndCreatedAt)
		if parseErr != nil {
			return nil, errors.Wrap(parseErr, "解析结束创建时间失败")
		}
		db = db.Where("created_at <= ?", endTime)
	}

	// 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询公告总数失败")
	}

	// 分页查询
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list []model.AnnouncementsInfo
	if err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "查询公告列表失败")
	}

	return &types.PageResult{
		List:     conv.ToTypesAnnouncementInfos(list),
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
