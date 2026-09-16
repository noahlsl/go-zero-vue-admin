package sysDictionary

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExportSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportSysDictionaryLogic {
	return &ExportSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ExportSysDictionary 导出字典 JSON（含字典详情）。
// 仅保留业务字段，剔除 ID、创建时间、更新时间等系统字段，与原 Gin 导出结构一致。
func (l *ExportSysDictionaryLogic) ExportSysDictionary(req *types.GetById) (resp *types.ExportSysDictionaryRes, err error) {
	if req.ID == 0 {
		return nil, errors.New("字典ID不能为空")
	}

	var dict model.SysDictionary
	if err = l.svcCtx.DB.WithContext(l.ctx).First(&dict, req.ID).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典失败")
	}

	var details []model.SysDictionaryDetail
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("sys_dictionary_id = ?", dict.ID).
		Order("sort").
		Find(&details).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典详情失败")
	}

	detailList := make([]types.ExportSysDictionaryDetail, 0, len(details))
	for _, d := range details {
		detailList = append(detailList, types.ExportSysDictionaryDetail{
			Label:  d.Label,
			Value:  d.Value,
			Extend: d.Extend,
			Status: d.Status,
			Sort:   d.Sort,
			Level:  d.Level,
			Path:   d.Path,
		})
	}

	return &types.ExportSysDictionaryRes{
		Name:                 dict.Name,
		Type:                 dict.Type,
		Status:               dict.Status,
		Desc:                 dict.Desc,
		SysDictionaryDetails: detailList,
	}, nil
}
