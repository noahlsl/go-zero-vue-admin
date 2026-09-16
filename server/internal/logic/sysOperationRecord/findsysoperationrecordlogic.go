package sysOperationRecord

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysOperationRecordLogic {
	return &FindSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysOperationRecordLogic) FindSysOperationRecord(req *types.GetById) (resp *types.SysOperationRecordRes, err error) {
	// 1. 根据ID查询操作记录
	var record model.SysOperationRecord
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&record).Error; err != nil {
		return nil, errors.Wrap(err, "查询操作记录失败")
	}

	return &types.SysOperationRecordRes{
		SysOperationRecord: modelOperationRecordToTypes(record),
	}, nil
}

// modelOperationRecordToTypes 将 model.SysOperationRecord 转换为 types.SysOperationRecord
func modelOperationRecordToTypes(m model.SysOperationRecord) types.SysOperationRecord {
	return types.SysOperationRecord{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		UserId:    uint(m.UserID),
		Ip:        m.Ip,
		Method:    m.Method,
		Path:      m.Path,
		Agent:     m.Agent,
		Status:    m.Status,
		Latency:   int64(m.Latency),
		Body:      m.Body,
		Resp:      m.Resp,
	}
}
