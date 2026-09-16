// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictionaryDetailListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictionaryDetailListLogic {
	return &GetSysDictionaryDetailListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictionaryDetailListLogic) GetSysDictionaryDetailList(req *types.GetSysDictionaryDetailListReq) (resp *types.PageResult, err error) {
	// todo: add your logic here and delete this line

	return
}
