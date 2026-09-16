package api

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiByIdLogic) GetApiById(req *types.GetById) (resp *types.SysApi, err error) {
	var apiModel model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).First(&apiModel, "id = ?", req.ID).Error; err != nil {
		return nil, errors.Wrap(err, "查询API失败")
	}

	return &types.SysApi{
		ID:          apiModel.ID,
		CreatedAt:   apiModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   apiModel.UpdatedAt.Format(time.RFC3339),
		Method:      apiModel.Method,
		Path:        apiModel.Path,
		ApiGroup:    apiModel.ApiGroup,
		Description: apiModel.Description,
	}, nil
}
