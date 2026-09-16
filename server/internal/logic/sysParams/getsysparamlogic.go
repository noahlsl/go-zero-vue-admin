package sysParams

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysParamLogic {
	return &GetSysParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysParamLogic) GetSysParam(req *types.GetById) (resp *types.Response, err error) {
	// 1. 根据ID查询单个参数
	var sysParams model.SysParams
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&sysParams).Error; err != nil {
		return nil, errors.Wrap(err, "查询参数失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "查询成功",
		Data: sysParams,
	}, nil
}
