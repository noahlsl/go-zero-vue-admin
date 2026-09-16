// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkillListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSkillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkillListLogic {
	return &GetSkillListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSkillListLogic) GetSkillList(req *types.GetSkillListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
