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

type GetDictionaryTreeListByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryTreeListByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryTreeListByTypeLogic {
	return &GetDictionaryTreeListByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetDictionaryTreeListByType 根据字典类型获取树形结构。
// 响应结构对齐 Gin 的 gin.H{"list": list}。
func (l *GetDictionaryTreeListByTypeLogic) GetDictionaryTreeListByType(req *types.GetDictionaryTreeListByTypeReq) (resp *types.DictionaryTreeRes, err error) {
	if req.Type == "" {
		return nil, errors.New("字典类型不能为空")
	}

	var detailList []model.SysDictionaryDetail
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysDictionaryDetail{}).
		Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id").
		Where("sys_dictionaries.type = ? AND sys_dictionary_details.parent_id IS NULL", req.Type).
		Order("sys_dictionary_details.sort").
		Find(&detailList).Error; err != nil {
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
func (l *GetDictionaryTreeListByTypeLogic) loadChildren(detail *model.SysDictionaryDetail) error {
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
