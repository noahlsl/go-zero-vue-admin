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

type GetSkillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSkillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkillDetailLogic {
	return &GetSkillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSkillDetailLogic) GetSkillDetail(req *types.GetById) (resp *types.SkillDetailRes, err error) {
	var skill model.SysSkill
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&skill).Error; err != nil {
		return nil, errors.Wrap(err, "查询技能详情失败")
	}

	return &types.SkillDetailRes{
		Skill: types.Skills{
			ID:        skill.ID,
			CreatedAt: skill.CreatedAt.Format(time.RFC3339),
			UpdatedAt: skill.UpdatedAt.Format(time.RFC3339),
			Name:      skill.Name,
			Desc:      skill.Desc,
			Content:   skill.Content,
		},
	}, nil
}
