// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateScriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateScriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateScriptLogic {
	return &CreateScriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateScriptLogic) CreateScript(req *types.CreateScriptReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
