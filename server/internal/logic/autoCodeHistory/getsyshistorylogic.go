package autoCodeHistory

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysHistoryLogic {
	return &GetSysHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysHistoryLogic) GetSysHistory(req *types.GetHistoryListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysAutoCodeHistory{})

	// 关键词搜索
	if req.Keyword != "" {
		db = db.Where("table_name LIKE ? OR struct_name LIKE ? OR abbreviation LIKE ?",
			"%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 查询总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询历史记录总数失败")
	}

	// 分页查询
	var histories []model.SysAutoCodeHistory
	offset := (req.Page - 1) * req.PageSize
	if offset < 0 {
		offset = 0
	}
	if err = db.Order("updated_at DESC").Offset(offset).Limit(req.PageSize).Find(&histories).Error; err != nil {
		return nil, errors.Wrap(err, "查询历史记录列表失败")
	}

	// 转换为响应类型
	list := make([]types.SysAutoCodeHistory, 0, len(histories))
	for _, h := range histories {
		list = append(list, types.SysAutoCodeHistory{
			ID:        h.ID,
			CreatedAt: h.CreatedAt.Format(time.RFC3339),
			UpdatedAt: h.UpdatedAt.Format(time.RFC3339),
			TableName: h.Table,
			Template:  h.Package,
			Author:    h.StructName,
			Migrate:   h.Flag == 0,
			Rollback:  h.Flag == 1,
		})
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
