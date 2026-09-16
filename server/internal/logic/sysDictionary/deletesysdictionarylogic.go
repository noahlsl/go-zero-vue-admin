package sysDictionary

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictionaryLogic {
	return &DeleteSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictionaryLogic) DeleteSysDictionary(req *types.GetById) (resp *types.Response, err error) {
	// 查询字典是否存在
	var dict model.SysDictionary
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&dict).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典失败")
	}

	// 删除关联的字典详情
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("sys_dictionary_id = ?", dict.ID).Delete(&model.SysDictionaryDetail{}).Error; err != nil {
		return nil, errors.Wrap(err, "删除字典详情失败")
	}

	// 删除字典本身
	if err = l.svcCtx.DB.WithContext(l.ctx).Delete(&dict).Error; err != nil {
		return nil, errors.Wrap(err, "删除字典失败")
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}
