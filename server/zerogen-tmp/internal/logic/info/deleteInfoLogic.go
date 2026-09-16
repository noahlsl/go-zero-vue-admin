// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInfoLogic {
	return &DeleteInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInfoLogic) DeleteInfo(req *types.DeleteAnnouncementInfoReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
