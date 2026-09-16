// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryPathLogic {
	return &GetDictionaryPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryPathLogic) GetDictionaryPath(req *types.GetByLowerId) (resp *types.DictionaryPathRes, err error) {
	// todo: add your logic here and delete this line

	return
}
