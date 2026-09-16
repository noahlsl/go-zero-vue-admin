package autoCode

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAIWorkflowSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAIWorkflowSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAIWorkflowSessionLogic {
	return &DeleteAIWorkflowSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAIWorkflowSessionLogic) DeleteAIWorkflowSession(req *types.DeleteAIWorkflowSessionReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("会话ID不能为空")
	}

	result := l.svcCtx.DB.WithContext(l.ctx).Delete(&model.SysAIWorkflowSession{}, req.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "删除会话失败")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("会话不存在")
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
