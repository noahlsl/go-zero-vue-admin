// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailsByParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailsByParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailsByParentLogic {
	return &GetDictionaryDetailsByParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryDetailsByParentLogic) GetDictionaryDetailsByParent(req *types.GetDictionaryDetailsByParentReq) (resp *types.DictionaryTreeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
