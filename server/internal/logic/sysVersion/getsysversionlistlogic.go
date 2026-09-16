package sysVersion

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysVersionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysVersionListLogic {
	return &GetSysVersionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysVersionListLogic) GetSysVersionList(req *types.GetSysVersionListReq) (resp *types.PageResult, err error) {
	var total int64
	var versions []model.SysVersion

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysVersion{})

	if err := db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "统计版本记录失败")
	}

	offset, limit := pageConvert(req.Page, req.PageSize)
	if err := db.Offset(offset).Limit(limit).
		Order("id DESC").
		Find(&versions).Error; err != nil {
		return nil, errors.Wrap(err, "查询版本列表失败")
	}

	list := make([]types.SysVersion, 0, len(versions))
	for _, v := range versions {
		list = append(list, types.SysVersion{
			ID:          v.ID,
			CreatedAt:   v.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			VersionName: v.VersionName,
			VersionCode: v.VersionCode,
			Description: v.Description,
			VersionData: v.VersionData,
		})
	}

	resp = &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return resp, nil
}

func pageConvert(page, pageSize int) (offset, limit int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return (page - 1) * pageSize, pageSize
}
