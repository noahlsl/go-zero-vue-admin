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

type GetAIWorkflowSessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAIWorkflowSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAIWorkflowSessionListLogic {
	return &GetAIWorkflowSessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAIWorkflowSessionListLogic) GetAIWorkflowSessionList(req *types.GetAIWorkflowSessionListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysAIWorkflowSession{})

	// 关键词搜索
	if req.Keyword != "" {
		db = db.Where("title LIKE ? OR summary LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 查询总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询会话总数失败")
	}

	// 分页查询
	var sessions []model.SysAIWorkflowSession
	offset := (req.Page - 1) * req.PageSize
	if offset < 0 {
		offset = 0
	}
	if err = db.Order("updated_at DESC").Offset(offset).Limit(req.PageSize).Find(&sessions).Error; err != nil {
		return nil, errors.Wrap(err, "查询会话列表失败")
	}

	// 转换为响应类型
	list := make([]types.AIWorkflowSession, 0, len(sessions))
	for _, s := range sessions {
		list = append(list, types.AIWorkflowSession{
			ID:        s.ID,
			CreatedAt: s.CreatedAt.Format(time.RFC3339),
			UpdatedAt: s.UpdatedAt.Format(time.RFC3339),
			Name:      s.Title,
			Content:   s.Summary,
			UserId:    s.UserID,
		})
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
