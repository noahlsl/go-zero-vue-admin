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

type GetTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplateLogic {
	return &GetTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateLogic) GetTemplate(req *types.GetTemplateReq) (resp *types.TemplateDetailRes, err error) {
	var tpl model.SysSkillTemplate
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&tpl).Error; err != nil {
		return nil, errors.Wrap(err, "查询模板失败")
	}

	return &types.TemplateDetailRes{
		Template: types.Templates{
			ID:        tpl.ID,
			CreatedAt: tpl.CreatedAt.Format(time.RFC3339),
			UpdatedAt: tpl.UpdatedAt.Format(time.RFC3339),
			Name:      tpl.Name,
			Content:   tpl.Content,
		},
	}, nil
}
