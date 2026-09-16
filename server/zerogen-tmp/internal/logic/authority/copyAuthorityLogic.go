// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authority

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyAuthorityLogic {
	return &CopyAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyAuthorityLogic) CopyAuthority(req *types.CopyAuthorityReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
