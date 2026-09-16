// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionary

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysDictionaryLogic {
	return &CreateSysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysDictionaryLogic) CreateSysDictionary(req *types.CreateSysDictionaryReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
