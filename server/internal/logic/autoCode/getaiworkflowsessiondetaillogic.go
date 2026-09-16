package autoCode

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAIWorkflowSessionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAIWorkflowSessionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAIWorkflowSessionDetailLogic {
	return &GetAIWorkflowSessionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAIWorkflowSessionDetailLogic) GetAIWorkflowSessionDetail(req *types.GetById) (resp *types.AIWorkflowSessionDetailRes, err error) {
	var session model.SysAIWorkflowSession
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&session).Error; err != nil {
		return nil, errors.Wrap(err, "查询会话详情失败")
	}

	return &types.AIWorkflowSessionDetailRes{
		Session: types.AIWorkflowSession{
			ID:        session.ID,
			CreatedAt: session.CreatedAt.Format(time.RFC3339),
			UpdatedAt: session.UpdatedAt.Format(time.RFC3339),
			Name:      session.Title,
			Content:   session.Summary,
			UserId:    session.UserID,
		},
	}, nil
}
