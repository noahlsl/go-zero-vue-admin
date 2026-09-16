package sysOperationRecord

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysOperationRecordLogic {
	return &DeleteSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysOperationRecordLogic) DeleteSysOperationRecord(req *types.DeleteSysOperationRecordReq) (resp *types.Response, err error) {
	// 1. 校验记录是否存在
	var record model.SysOperationRecord
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&record).Error; err != nil {
		return nil, errors.Wrap(err, "操作记录不存在")
	}

	// 2. 执行软删除
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).Delete(&model.SysOperationRecord{}).Error; err != nil {
		l.Logger.Errorw("删除操作记录失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysOperationRecord"),
			logx.Field("action", "delete"),
		)
		return nil, errors.Wrap(err, "删除操作记录失败")
	}

	return &types.Response{Msg: "删除成功"}, nil
}
