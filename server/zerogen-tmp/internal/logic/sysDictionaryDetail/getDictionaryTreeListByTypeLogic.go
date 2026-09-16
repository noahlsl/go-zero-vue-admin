// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryTreeListByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryTreeListByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryTreeListByTypeLogic {
	return &GetDictionaryTreeListByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryTreeListByTypeLogic) GetDictionaryTreeListByType(req *types.GetDictionaryTreeListByTypeReq) (resp *types.DictionaryTreeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
