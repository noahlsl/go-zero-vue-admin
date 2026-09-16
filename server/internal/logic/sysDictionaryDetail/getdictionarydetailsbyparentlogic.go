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

type GetDictionaryDetailsByParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailsByParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailsByParentLogic {
	return &GetDictionaryDetailsByParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetDictionaryDetailsByParent 根据父级ID获取字典详情。
// 响应结构对齐 Gin 的 gin.H{"list": list}，list 为完整 SysDictionaryDetail 数组。
func (l *GetDictionaryDetailsByParentLogic) GetDictionaryDetailsByParent(req *types.GetDictionaryDetailsByParentReq) (resp *types.DictionaryTreeRes, err error) {
	if req.SysDictionaryID == 0 {
		return nil, errors.New("字典ID不能为空")
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysDictionaryDetail{}).
		Where("sys_dictionary_id = ?", req.SysDictionaryID)

	if req.ParentID != nil {
		db = db.Where("parent_id = ?", *req.ParentID)
	} else {
		db = db.Where("parent_id IS NULL")
	}

	var detailList []model.SysDictionaryDetail
	if err = db.Order("sort").Find(&detailList).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典详情失败")
	}

	for i := range detailList {
		setDisabled(&detailList[i])
	}

	// 与原 Gin 行为一致：仅当请求显式要求时才递归加载子级
	if req.IncludeChildren {
		for i := range detailList {
			if err = l.loadChildren(&detailList[i]); err != nil {
				return nil, err
			}
		}
	}

	return &types.DictionaryTreeRes{List: conv.ToTypesSysDictionaryDetails(detailList)}, nil
}

// loadChildren 递归加载子级并计算 disabled
func (l *GetDictionaryDetailsByParentLogic) loadChildren(detail *model.SysDictionaryDetail) error {
	var children []model.SysDictionaryDetail
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("parent_id = ?", detail.ID).Order("sort").Find(&children).Error; err != nil {
		return errors.Wrap(err, "查询字典子级失败")
	}

	for i := range children {
		setDisabled(&children[i])
		if err := l.loadChildren(&children[i]); err != nil {
			return err
		}
	}

	detail.Children = children
	return nil
}
