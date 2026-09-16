// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveGlobalConstraintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveGlobalConstraintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveGlobalConstraintLogic {
	return &SaveGlobalConstraintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveGlobalConstraintLogic) SaveGlobalConstraint(req *types.SaveGlobalConstraintReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
