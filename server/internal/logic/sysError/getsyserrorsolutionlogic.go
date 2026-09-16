package sysError

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysErrorSolutionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysErrorSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysErrorSolutionLogic {
	return &GetSysErrorSolutionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysErrorSolutionLogic) GetSysErrorSolution(req *types.GetByLowerId) (resp *types.SysErrorRes, err error) {
	// 1. 查询错误日志
	var sysError model.SysError
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&sysError).Error; err != nil {
		return nil, errors.Wrap(err, "查询错误日志失败")
	}

	// 2. 更新状态为处理中
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysError{}).Where("id = ?", req.ID).
		Update("status", "处理中").Error; err != nil {
		return nil, errors.Wrap(err, "更新错误状态失败")
	}

	// 3. 异步获取解决方案（更新状态为已处理）
	go l.fetchSolution(sysError)

	return &types.SysErrorRes{
		ErrorInfo: modelSysErrorToTypes(sysError),
	}, nil
}

// fetchSolution 异步获取错误解决方案并更新状态
func (l *GetSysErrorSolutionLogic) fetchSolution(sysError model.SysError) {
	solution := "系统已自动记录错误信息，请人工排查处理"
	status := "已处理"

	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysError{}).Where("id = ?", sysError.ID).
		Updates(map[string]interface{}{
			"solution": solution,
			"status":   status,
		}).Error; err != nil {
		l.Logger.Errorw("更新错误解决方案失败",
			logx.Field("error", err),
			logx.Field("id", sysError.ID),
			logx.Field("module", "sysError"),
			logx.Field("action", "fetch_solution"),
		)
	}
}
