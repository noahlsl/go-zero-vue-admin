// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInfoLogic {
	return &CreateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInfoLogic) CreateInfo(req *types.CreateAnnouncementInfoReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
