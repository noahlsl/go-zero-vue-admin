// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysDictionaryLogic {
	return &FindSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysDictionaryLogic) FindSysDictionary(req *types.FindSysDictionaryReq) (resp *types.SysDictionaryRes, err error) {
	// todo: add your logic here and delete this line

	return
}
