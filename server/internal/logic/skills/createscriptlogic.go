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

type CreateScriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateScriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateScriptLogic {
	return &CreateScriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateScriptLogic) CreateScript(req *types.CreateScriptReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("脚本名不能为空")
	}

	script := model.SysSkillScript{
		Name:    req.Name,
		Content: req.Content,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&script).Error; err != nil {
		return nil, errors.Wrap(err, "创建脚本失败")
	}

	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
