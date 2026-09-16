package autoCode

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAIWorkflowSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveAIWorkflowSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAIWorkflowSessionLogic {
	return &SaveAIWorkflowSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAIWorkflowSessionLogic) SaveAIWorkflowSession(req *types.SaveAIWorkflowSessionReq) (resp *types.Response, err error) {
	session := model.SysAIWorkflowSession{
		Title:   req.Name,
		Summary: req.Content,
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&session).Error; err != nil {
		return nil, errors.Wrap(err, "保存 AI 工作流会话失败")
	}

	return &types.Response{
		Code: 0,
		Data: session,
		Msg:  "保存成功",
	}, nil
}
