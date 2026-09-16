package skills

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveGlobalConstraintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveGlobalConstraintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveGlobalConstraintLogic {
	return &SaveGlobalConstraintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveGlobalConstraintLogic) SaveGlobalConstraint(req *types.SaveGlobalConstraintReq) (resp *types.Response, err error) {
	// 查询现有约束
	var constraint model.SysSkillGlobalConstraint
	err = l.svcCtx.DB.WithContext(l.ctx).Order("id DESC").First(&constraint).Error

	if err != nil {
		// 没有记录则创建
		constraint = model.SysSkillGlobalConstraint{
			Content: req.Content,
		}
		if err = l.svcCtx.DB.WithContext(l.ctx).Create(&constraint).Error; err != nil {
			return nil, errors.Wrap(err, "创建全局约束失败")
		}
	} else {
		// 更新
		constraint.Content = req.Content
		if err = l.svcCtx.DB.WithContext(l.ctx).Save(&constraint).Error; err != nil {
			return nil, errors.Wrap(err, "更新全局约束失败")
		}
	}

	return &types.Response{Code: 0, Msg: "保存成功"}, nil
}
