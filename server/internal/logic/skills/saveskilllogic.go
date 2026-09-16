package skills

import (
	"context"
	"strings"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSkillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSkillLogic {
	return &SaveSkillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveSkillLogic) SaveSkill(req *types.SaveSkillReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("技能名称不能为空")
	}

	if req.ID != 0 {
		// 更新
		var skill model.SysSkill
		if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&skill).Error; err != nil {
			return nil, errors.Wrap(err, "技能不存在")
		}
		skill.Name = req.Name
		skill.Desc = req.Desc
		skill.Content = req.Content
		if err = l.svcCtx.DB.WithContext(l.ctx).Save(&skill).Error; err != nil {
			return nil, errors.Wrap(err, "更新技能失败")
		}
		return &types.Response{Code: 0, Msg: "更新成功"}, nil
	}

	// 创建
	skill := model.SysSkill{
		Name:    req.Name,
		Desc:    req.Desc,
		Content: req.Content,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&skill).Error; err != nil {
		return nil, errors.Wrap(err, "创建技能失败")
	}
	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
