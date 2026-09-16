// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package sysDictionaryDetail

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictionaryDetailLogic {
	return &DeleteSysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictionaryDetailLogic) DeleteSysDictionaryDetail(req *types.DeleteSysDictionaryDetailReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
