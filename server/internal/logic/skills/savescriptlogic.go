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

type SaveScriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveScriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveScriptLogic {
	return &SaveScriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveScriptLogic) SaveScript(req *types.SaveScriptReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("脚本名不能为空")
	}

	if req.ID != 0 {
		var script model.SysSkillScript
		if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&script).Error; err != nil {
			return nil, errors.Wrap(err, "脚本不存在")
		}
		script.Name = req.Name
		script.Content = req.Content
		if err = l.svcCtx.DB.WithContext(l.ctx).Save(&script).Error; err != nil {
			return nil, errors.Wrap(err, "更新脚本失败")
		}
		return &types.Response{Code: 0, Msg: "更新成功"}, nil
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
