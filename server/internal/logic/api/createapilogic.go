package api

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiLogic) CreateApi(req *types.CreateApiReq) (resp *types.Response, err error) {
	// 检查是否存在相同 path + method 的 API
	var count int64
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysApi{}).
		Where("path = ? AND method = ?", req.Path, req.Method).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询API是否存在失败")
	}
	if count > 0 {
		return nil, errors.New("存在相同api")
	}

	apiModel := model.SysApi{
		Path:        req.Path,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
		Method:      req.Method,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&apiModel).Error; err != nil {
		return nil, errors.Wrap(err, "创建API失败")
	}

	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
