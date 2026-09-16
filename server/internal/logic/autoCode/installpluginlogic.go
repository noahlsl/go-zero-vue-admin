package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
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

// InstallPlugin 安装插件
// TODO: 实现完整的插件安装逻辑（解压、注入代码等）
func (l *InstallPluginLogic) InstallPlugin(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("插件ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Data: map[string]interface{}{
			"web":    0,
			"server": 0,
		},
		Msg: fmt.Sprintf("插件 %d 安装成功", req.ID),
	}, nil
}
