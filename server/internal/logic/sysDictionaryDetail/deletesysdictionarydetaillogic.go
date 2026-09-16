package sysDictionaryDetail

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictionaryDetailLogic {
	return &DeleteSysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictionaryDetailLogic) DeleteSysDictionaryDetail(req *types.DeleteSysDictionaryDetailReq) (resp *types.Response, err error) {
	// 检查是否有子项
	var count int64
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysDictionaryDetail{}).
		Where("parent_id = ?", req.ID).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询子项数量失败")
	}
	if count > 0 {
		return nil, errors.New("该字典详情下还有子项，无法删除")
	}

	// 删除字典详情
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("id = ?", req.ID).Delete(&model.SysDictionaryDetail{}).Error; err != nil {
		return nil, errors.Wrap(err, "删除字典详情失败")
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}
