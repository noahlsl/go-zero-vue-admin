package sysError

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysErrorLogic {
	return &CreateSysErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysErrorLogic) CreateSysError(req *types.CreateSysErrorReq) (resp *types.Response, err error) {
	// 1. 构建错误日志记录
	sysError := model.SysError{
		Form:     req.Form,
		Info:     req.Info,
		Level:    req.Level,
		Solution: req.Solution,
		Status:   "未处理",
	}

	// 2. 写入数据库
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&sysError).Error; err != nil {
		l.Logger.Errorw("创建错误日志失败",
			logx.Field("error", err),
			logx.Field("module", "sysError"),
			logx.Field("action", "create"),
		)
		return nil, errors.Wrap(err, "创建错误日志失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "创建成功",
		Data: sysError,
	}, nil
}
