// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package info

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoListLogic {
	return &GetInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoListLogic) GetInfoList(req *types.GetAnnouncementInfoListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
