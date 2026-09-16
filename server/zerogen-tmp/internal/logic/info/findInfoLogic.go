// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindInfoLogic {
	return &FindInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindInfoLogic) FindInfo(req *types.FindAnnouncementInfoReq) (resp *types.AnnouncementInfoRes, err error) {
	// todo: add your logic here and delete this line

	return
}
