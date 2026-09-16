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

type FindSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysDictionaryLogic {
	return &FindSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FindSysDictionary 按 type 或 id 查询字典单条数据。
// 查询条件对齐原 Gin：(type = ? OR id = ?) and status = ?，status 为空时按 true 处理。
// 响应 key 对齐 Gin 的 gin.H{"resysDictionary": ...}。
func (l *FindSysDictionaryLogic) FindSysDictionary(req *types.FindSysDictionaryReq) (resp *types.SysDictionaryRes, err error) {
	status := true
	if req.Status != nil {
		status = *req.Status
	}

	var dict model.SysDictionary
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("(type = ? OR id = ?) and status = ?", req.Type, req.ID, status).
		First(&dict).Error; err != nil {
		return nil, errors.Wrap(err, "字典未创建或未开启")
	}

	return &types.SysDictionaryRes{
		ResysDictionary: conv.ToTypesSysDictionary(dict),
	}, nil
}
