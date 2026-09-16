// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInfoByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInfoByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInfoByIdsLogic {
	return &DeleteInfoByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInfoByIdsLogic) DeleteInfoByIds(req *types.DeleteAnnouncementInfoByIdsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
