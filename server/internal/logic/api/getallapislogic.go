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

type GetAllApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllApisLogic {
	return &GetAllApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllApisLogic) GetAllApis(req *types.GetAllApisReq) (resp []types.SysApi, err error) {
	var apiList []model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).Order("id desc").Find(&apiList).Error; err != nil {
		return nil, errors.Wrap(err, "查询所有API失败")
	}

	resp = make([]types.SysApi, 0, len(apiList))
	for _, item := range apiList {
		resp = append(resp, types.SysApi{
			ID:          item.ID,
			CreatedAt:   item.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
			Method:      item.Method,
			Path:        item.Path,
			ApiGroup:    item.ApiGroup,
			Description: item.Description,
		})
	}

	return resp, nil
}
