package autoCodeHistory

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysHistoryLogic {
	return &DelSysHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysHistoryLogic) DelSysHistory(req *types.DeleteHistoryReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	result := l.svcCtx.DB.WithContext(l.ctx).Delete(&model.SysAutoCodeHistory{}, req.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "删除历史记录失败")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("历史记录不存在")
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
