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

type CreateTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTemplateLogic {
	return &CreateTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTemplateLogic) CreateTemplate(req *types.CreateTemplateReq) (resp *types.Response, err error) {
	name := req.Name
	if strings.TrimSpace(name) == "" {
		name = req.TableName
	}
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("模板名称不能为空")
	}

	tpl := model.SysSkillTemplate{
		Name:    name,
		Content: req.DataType,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&tpl).Error; err != nil {
		return nil, errors.Wrap(err, "创建模板失败")
	}

	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
