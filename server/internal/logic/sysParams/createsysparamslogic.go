package sysParams

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysParamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysParamsLogic {
	return &CreateSysParamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysParamsLogic) CreateSysParams(req *types.CreateSysParamsReq) (resp *types.Response, err error) {
	// 1. 参数校验
	if req.Key == "" {
		return nil, errors.New("参数键不能为空")
	}

	// 2. 校验 key 是否已存在
	var exist model.SysParams
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("key = ?", req.Key).First(&exist).Error; err == nil {
		return nil, errors.New("参数键已存在")
	}

	// 3. 创建参数记录
	sysParams := model.SysParams{
		Name:  req.Name,
		Key:   req.Key,
		Value: req.Value,
		Desc:  req.Desc,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&sysParams).Error; err != nil {
		l.Logger.Errorw("创建参数失败",
			logx.Field("error", err),
			logx.Field("key", req.Key),
			logx.Field("module", "sysParams"),
			logx.Field("action", "create"),
		)
		return nil, errors.Wrap(err, "创建参数失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "创建成功",
		Data: sysParams,
	}, nil
}
