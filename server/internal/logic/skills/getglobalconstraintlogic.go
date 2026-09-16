package skills

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGlobalConstraintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGlobalConstraintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGlobalConstraintLogic {
	return &GetGlobalConstraintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGlobalConstraintLogic) GetGlobalConstraint() (resp *types.GlobalConstraintRes, err error) {
	var constraint model.SysSkillGlobalConstraint
	// 获取最新的全局约束（取 ID 最大的记录）
	if err = l.svcCtx.DB.WithContext(l.ctx).Order("id DESC").First(&constraint).Error; err != nil {
		// 如果没有记录，返回默认内容
		return &types.GlobalConstraintRes{
			Constraint: types.GlobalConstraint{
				Content: "# 全局约束\n请在这里补充该工具的统一约束与使用规范。\n",
			},
		}, nil
	}

	return &types.GlobalConstraintRes{
		Constraint: types.GlobalConstraint{
			ID:        constraint.ID,
			CreatedAt: constraint.CreatedAt.Format(time.RFC3339),
			UpdatedAt: constraint.UpdatedAt.Format(time.RFC3339),
			Content:   constraint.Content,
		},
	}, nil
}
