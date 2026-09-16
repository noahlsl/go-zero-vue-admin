package sysDictionaryDetail

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictionaryDetailListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictionaryDetailListLogic {
	return &GetSysDictionaryDetailListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictionaryDetailListLogic) GetSysDictionaryDetailList(req *types.GetSysDictionaryDetailListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysDictionaryDetail{})

	// 条件过滤
	if req.Label != "" {
		db = db.Where("label LIKE ?", "%"+req.Label+"%")
	}
	if req.Value != "" {
		db = db.Where("value = ?", req.Value)
	}
	if req.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", req.SysDictionaryID)
	}

	// 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "统计字典详情总数失败")
	}

	// 分页查询
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var detailList []model.SysDictionaryDetail
	if err = db.Limit(limit).Offset(offset).Order("sort").Order("id").Find(&detailList).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典详情列表失败")
	}

	return &types.PageResult{
		List:     conv.ToTypesSysDictionaryDetails(detailList),
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
