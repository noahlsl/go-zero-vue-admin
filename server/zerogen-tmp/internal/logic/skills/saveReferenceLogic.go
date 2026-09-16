// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveReferenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveReferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveReferenceLogic {
	return &SaveReferenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveReferenceLogic) SaveReference(req *types.SaveReferenceReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
