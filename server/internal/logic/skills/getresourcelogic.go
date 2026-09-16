package skills

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResourceLogic {
	return &GetResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetResourceLogic) GetResource(req *types.GetResourceReq) (resp *types.ResourceDetailRes, err error) {
	var res model.SysSkillResource
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&res).Error; err != nil {
		return nil, errors.Wrap(err, "查询资源文件失败")
	}

	return &types.ResourceDetailRes{
		Resource: types.Resources{
			ID:        res.ID,
			CreatedAt: res.CreatedAt.Format(time.RFC3339),
			UpdatedAt: res.UpdatedAt.Format(time.RFC3339),
			Name:      res.Name,
			Type:      res.Type,
			Content:   res.Content,
		},
	}, nil
}
