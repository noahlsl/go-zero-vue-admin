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

type SaveTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTemplateLogic {
	return &SaveTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveTemplateLogic) SaveTemplate(req *types.SaveTemplateReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("模板名称不能为空")
	}

	if req.ID != 0 {
		var tpl model.SysSkillTemplate
		if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&tpl).Error; err != nil {
			return nil, errors.Wrap(err, "模板不存在")
		}
		tpl.Name = req.Name
		tpl.Content = req.Content
		if err = l.svcCtx.DB.WithContext(l.ctx).Save(&tpl).Error; err != nil {
			return nil, errors.Wrap(err, "更新模板失败")
		}
		return &types.Response{Code: 0, Msg: "更新成功"}, nil
	}

	tpl := model.SysSkillTemplate{
		Name:    req.Name,
		Content: req.Content,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&tpl).Error; err != nil {
		return nil, errors.Wrap(err, "创建模板失败")
	}
	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
