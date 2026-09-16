// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveScriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveScriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveScriptLogic {
	return &SaveScriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveScriptLogic) SaveScript(req *types.SaveScriptReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
