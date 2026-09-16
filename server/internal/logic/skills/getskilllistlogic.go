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

type GetSkillListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSkillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkillListLogic {
	return &GetSkillListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSkillListLogic) GetSkillList(req *types.GetSkillListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysSkill{})

	if req.Keyword != "" {
		db = db.Where("name LIKE ? OR desc LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询技能总数失败")
	}

	var skillList []model.SysSkill
	offset := (req.Page - 1) * req.PageSize
	if offset < 0 {
		offset = 0
	}
	if err = db.Order("updated_at DESC").Offset(offset).Limit(req.PageSize).Find(&skillList).Error; err != nil {
		return nil, errors.Wrap(err, "查询技能列表失败")
	}

	list := make([]types.Skills, 0, len(skillList))
	for _, s := range skillList {
		list = append(list, types.Skills{
			ID:        s.ID,
			CreatedAt: s.CreatedAt.Format(time.RFC3339),
			UpdatedAt: s.UpdatedAt.Format(time.RFC3339),
			Name:      s.Name,
			Desc:      s.Desc,
			Content:   s.Content,
		})
	}

	return &types.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
