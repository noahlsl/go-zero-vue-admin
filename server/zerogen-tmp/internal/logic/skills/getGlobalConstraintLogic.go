// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGlobalConstraintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGlobalConstraintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGlobalConstraintLogic {
	return &GetGlobalConstraintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGlobalConstraintLogic) GetGlobalConstraint() (resp *types.GlobalConstraintRes, err error) {
	// todo: add your logic here and delete this line

	return
}
