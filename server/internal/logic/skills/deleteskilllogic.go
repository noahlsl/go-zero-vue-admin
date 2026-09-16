package skills

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSkillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkillLogic {
	return &DeleteSkillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSkillLogic) DeleteSkill(req *types.DeleteSkillReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("技能ID不能为空")
	}

	result := l.svcCtx.DB.WithContext(l.ctx).Delete(&model.SysSkill{}, req.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "删除技能失败")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("技能不存在")
	}

	return &types.Response{Code: 0, Msg: "删除成功"}, nil
}
