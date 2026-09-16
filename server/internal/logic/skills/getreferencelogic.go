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

type GetReferenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReferenceLogic {
	return &GetReferenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReferenceLogic) GetReference(req *types.GetReferenceReq) (resp *types.ReferenceDetailRes, err error) {
	var ref model.SysSkillReference
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&ref).Error; err != nil {
		return nil, errors.Wrap(err, "查询参考资料失败")
	}

	return &types.ReferenceDetailRes{
		Reference: types.References{
			ID:        ref.ID,
			CreatedAt: ref.CreatedAt.Format(time.RFC3339),
			UpdatedAt: ref.UpdatedAt.Format(time.RFC3339),
			Name:      ref.Name,
			Content:   ref.Content,
		},
	}, nil
}
