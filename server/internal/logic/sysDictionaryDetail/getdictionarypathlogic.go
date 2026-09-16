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

type GetDictionaryPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryPathLogic {
	return &GetDictionaryPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetDictionaryPath 获取字典详情的完整层级路径。
// 响应结构对齐 Gin 的 gin.H{"path": path}，path 为完整 SysDictionaryDetail 数组。
func (l *GetDictionaryPathLogic) GetDictionaryPath(req *types.GetByLowerId) (resp *types.DictionaryPathRes, err error) {
	if req.ID == 0 {
		return nil, errors.New("字典详情ID不能为空")
	}

	path, err := l.getPathByID(uint(req.ID))
	if err != nil {
		return nil, err
	}

	return &types.DictionaryPathRes{Path: conv.ToTypesSysDictionaryDetails(path)}, nil
}

// getPathByID 递归向上查找父级，返回从根到当前节点的完整路径
func (l *GetDictionaryPathLogic) getPathByID(id uint) ([]model.SysDictionaryDetail, error) {
	var detail model.SysDictionaryDetail
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&detail, id).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典详情失败")
	}

	path := []model.SysDictionaryDetail{detail}
	if detail.ParentID == nil {
		return path, nil
	}

	parentPath, err := l.getPathByID(*detail.ParentID)
	if err != nil {
		return nil, err
	}
	return append(parentPath, path...), nil
}
