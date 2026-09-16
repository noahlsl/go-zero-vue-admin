// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateReferenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReferenceLogic {
	return &CreateReferenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReferenceLogic) CreateReference(req *types.CreateReferenceReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
