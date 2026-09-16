package file

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileLogic {
	return &FindFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFileLogic) FindFile(req *types.GetById) (resp *types.ExaFileRes, err error) {
	// 1. 根据 ID 查询文件记录
	var file model.ExaFileUploadAndDownload
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&file).Error; err != nil {
		return nil, errors.Wrap(err, "查询文件记录失败")
	}

	// 2. 转换为响应类型
	return &types.ExaFileRes{
		File: l.toResponse(file),
	}, nil
}

func (l *FindFileLogic) toResponse(m model.ExaFileUploadAndDownload) types.ExaFileUploadAndDownload {
	return types.ExaFileUploadAndDownload{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		Name:      m.Name,
		Url:       m.Url,
	}
}
