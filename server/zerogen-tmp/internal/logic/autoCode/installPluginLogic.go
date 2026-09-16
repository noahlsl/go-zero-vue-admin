// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package autoCode

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InstallPluginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInstallPluginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InstallPluginLogic {
	return &InstallPluginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InstallPluginLogic) InstallPlugin(req *types.DeletePackageReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
