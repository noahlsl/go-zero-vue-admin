// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysDictionaryDetailLogic {
	return &FindSysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysDictionaryDetailLogic) FindSysDictionaryDetail(req *types.GetById) (resp *types.SysDictionaryDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
