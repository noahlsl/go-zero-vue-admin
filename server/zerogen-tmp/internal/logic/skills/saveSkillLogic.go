// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
