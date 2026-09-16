package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePluginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePluginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePluginLogic {
	return &RemovePluginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RemovePlugin 移除插件
// TODO: 实现完整的插件移除逻辑
func (l *RemovePluginLogic) RemovePlugin(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("插件ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Msg:  fmt.Sprintf("插件 %d 移除成功", req.ID),
	}, nil
}
