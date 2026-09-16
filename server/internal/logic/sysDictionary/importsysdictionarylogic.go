package sysDictionary

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ImportSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportSysDictionaryLogic {
	return &ImportSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportSysDictionaryLogic) ImportSysDictionary(req *types.ImportSysDictionaryReq) error {
	// 解析导入的 JSON
	var importData types.ImportSysDictionaryData
	if err := sonic.UnmarshalString(req.Json, &importData); err != nil {
		return errors.Wrap(err, "JSON格式错误")
	}

	// 验证必填字段
	if importData.Name == "" {
		return errors.New("字典名称不能为空")
	}
	if importData.Type == "" {
		return errors.New("字典类型不能为空")
	}

	// 检查字典类型是否已存在
	var count int64
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysDictionary{}).
		Where("type = ?", importData.Type).
		Count(&count).Error; err != nil {
		return errors.Wrap(err, "查询字典类型是否存在失败")
	}
	if count > 0 {
		return errors.New("存在相同的type，不允许导入")
	}

	// 在事务中创建字典和详情
	return l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		dictionary := model.SysDictionary{
			Name:   importData.Name,
			Type:   importData.Type,
			Status: importData.Status,
			Desc:   importData.Desc,
		}
		if err := tx.Create(&dictionary).Error; err != nil {
			return errors.Wrap(err, "创建字典失败")
		}

		if len(importData.SysDictionaryDetails) == 0 {
			return nil
		}

		// 旧ID -> 新ID 映射，用于第二遍还原父子关系
		idMap := make(map[uint]uint)

		// 第一遍：创建所有详情记录
		for _, detail := range importData.SysDictionaryDetails {
			if detail.Label == "" || detail.Value == "" {
				continue
			}

			detailRecord := model.SysDictionaryDetail{
				Label:           detail.Label,
				Value:           detail.Value,
				Extend:          detail.Extend,
				Status:          detail.Status,
				Sort:            detail.Sort,
				Level:           detail.Level,
				Path:            detail.Path,
				SysDictionaryID: int(dictionary.ID),
			}
			if err := tx.Create(&detailRecord).Error; err != nil {
				return errors.Wrap(err, "创建字典详情失败")
			}
			if detail.ID > 0 {
				idMap[detail.ID] = detailRecord.ID
			}
		}

		// 第二遍：还原 parent_id 关系
		for _, detail := range importData.SysDictionaryDetails {
			if detail.ParentID == nil || *detail.ParentID == 0 || detail.ID == 0 {
				continue
			}
			newID, exists := idMap[detail.ID]
			if !exists {
				continue
			}
			newParentID, parentExists := idMap[*detail.ParentID]
			if !parentExists {
				continue
			}
			if err := tx.Model(&model.SysDictionaryDetail{}).
				Where("id = ?", newID).
				Update("parent_id", newParentID).Error; err != nil {
				return errors.Wrap(err, "还原字典详情父子关系失败")
			}
		}

		return nil
	})
}
