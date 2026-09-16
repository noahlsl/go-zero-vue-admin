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

type FindSysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysDictionaryDetailLogic {
	return &FindSysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FindSysDictionaryDetail 用 id 查询单条字典详情。
// 响应 key 对齐 Gin 的 gin.H{"reSysDictionaryDetail": ...}。
func (l *FindSysDictionaryDetailLogic) FindSysDictionaryDetail(req *types.GetById) (resp *types.SysDictionaryDetailRes, err error) {
	if req.ID == 0 {
		return nil, errors.New("id不能为空")
	}

	var detail model.SysDictionaryDetail
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("id = ?", req.ID).First(&detail).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典详情失败")
	}

	return &types.SysDictionaryDetailRes{
		ReSysDictionaryDetail: conv.ToTypesSysDictionaryDetail(detail),
	}, nil
}
