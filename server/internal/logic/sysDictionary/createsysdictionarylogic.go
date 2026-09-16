package sysDictionary

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysDictionaryLogic {
	return &CreateSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysDictionaryLogic) CreateSysDictionary(req *types.CreateSysDictionaryReq) (resp *types.Response, err error) {
	// 检查是否存在相同的 type
	var count int64
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysDictionary{}).
		Where("type = ?", req.Type).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典类型是否存在失败")
	}
	if count > 0 {
		return nil, errors.New("存在相同的type，不允许创建")
	}

	dict := model.SysDictionary{
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Desc:   req.Desc,
		ParentID: req.ParentID,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&dict).Error; err != nil {
		return nil, errors.Wrap(err, "创建字典失败")
	}

	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
