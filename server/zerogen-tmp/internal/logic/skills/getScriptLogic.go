// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScriptLogic {
	return &GetScriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScriptLogic) GetScript(req *types.GetScriptReq) (resp *types.ScriptDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
