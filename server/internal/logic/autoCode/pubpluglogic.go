package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PubPlugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPubPlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubPlugLogic {
	return &PubPlugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PubPlug 发布插件
// TODO: 实现完整的插件发布逻辑
func (l *PubPlugLogic) PubPlug(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("插件ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Msg:  fmt.Sprintf("插件 %d 发布成功", req.ID),
	}, nil
}
