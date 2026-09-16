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

type GetScriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScriptLogic {
	return &GetScriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScriptLogic) GetScript(req *types.GetScriptReq) (resp *types.ScriptDetailRes, err error) {
	var script model.SysSkillScript
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&script).Error; err != nil {
		return nil, errors.Wrap(err, "查询脚本失败")
	}

	return &types.ScriptDetailRes{
		Script: types.Scripts{
			ID:        script.ID,
			CreatedAt: script.CreatedAt.Format(time.RFC3339),
			UpdatedAt: script.UpdatedAt.Format(time.RFC3339),
			Name:      script.Name,
			Content:   script.Content,
		},
	}, nil
}
