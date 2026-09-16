// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package email

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailTestLogic {
	return &EmailTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailTestLogic) EmailTest() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
