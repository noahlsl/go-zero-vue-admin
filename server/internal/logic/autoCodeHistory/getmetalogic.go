package autoCodeHistory

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMetaLogic {
	return &GetMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMetaLogic) GetMeta(req *types.CreateHistoryReq) (resp *types.SysAutoCodeHistory, err error) {
	if req.ID == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	var history model.SysAutoCodeHistory
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&history).Error; err != nil {
		return nil, errors.Wrap(err, "查询历史记录失败")
	}

	return &types.SysAutoCodeHistory{
		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		TableName: history.Table,
		Template:  history.Package,
		Author:    history.StructName,
		Voice:     history.Request,
		Migrate:   history.Flag == 0,
		Rollback:  history.Flag == 1,
	}, nil
}
