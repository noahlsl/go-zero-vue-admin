// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictionaryListLogic {
	return &GetSysDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictionaryListLogic) GetSysDictionaryList(req *types.GetSysDictionaryListReq) (resp []types.SysDictionary, err error) {
	// todo: add your logic here and delete this line

	return
}
