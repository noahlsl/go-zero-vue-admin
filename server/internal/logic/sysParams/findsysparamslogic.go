package sysParams

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysParamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysParamsLogic {
	return &FindSysParamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysParamsLogic) FindSysParams(req *types.GetById) (resp *types.Response, err error) {
	// 1. 根据ID查询参数
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
