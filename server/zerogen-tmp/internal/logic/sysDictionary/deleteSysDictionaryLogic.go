// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictionaryLogic {
	return &DeleteSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictionaryLogic) DeleteSysDictionary(req *types.DeleteSysDictionaryReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
