// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReferenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReferenceLogic {
	return &GetReferenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReferenceLogic) GetReference(req *types.GetReferenceReq) (resp *types.ReferenceDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
