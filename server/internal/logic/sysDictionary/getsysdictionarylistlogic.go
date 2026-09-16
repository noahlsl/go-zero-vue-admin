package sysDictionary

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictionaryListLogic {
	return &GetSysDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetSysDictionaryList 获取字典列表。
// 响应与原 Gin 一致，直接返回字典数组（非分页结构），并预加载子字典。
func (l *GetSysDictionaryListLogic) GetSysDictionaryList(req *types.GetSysDictionaryListReq) (resp []types.SysDictionary, err error) {
	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysDictionary{})
	if req.Name != "" {
		db = db.Where("name LIKE ? OR type LIKE ?", "%"+req.Name+"%", "%"+req.Name+"%")
	}

	var dictList []model.SysDictionary
	if err = db.Preload("Children").Find(&dictList).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典列表失败")
	}

	return conv.ToTypesSysDictionaries(dictList), nil
}
