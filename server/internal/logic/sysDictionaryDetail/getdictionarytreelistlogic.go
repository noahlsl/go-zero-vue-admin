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

type GetDictionaryTreeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryTreeListLogic {
	return &GetDictionaryTreeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetDictionaryTreeList 获取指定字典的树形结构列表。
// 响应结构对齐 Gin 的 gin.H{"list": list}。
func (l *GetDictionaryTreeListLogic) GetDictionaryTreeList(req *types.GetDictionaryTreeListReq) (resp *types.DictionaryTreeRes, err error) {
	var detailList []model.SysDictionaryDetail
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("sys_dictionary_id = ? AND parent_id IS NULL", req.SysDictionaryID).
		Order("sort").Find(&detailList).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典树形列表失败")
	}

	for i := range detailList {
		setDisabled(&detailList[i])
		if err = l.loadChildren(&detailList[i]); err != nil {
			return nil, err
		}
	}

	return &types.DictionaryTreeRes{List: conv.ToTypesSysDictionaryDetails(detailList)}, nil
}

// loadChildren 递归加载子级并计算 disabled
func (l *GetDictionaryTreeListLogic) loadChildren(detail *model.SysDictionaryDetail) error {
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
