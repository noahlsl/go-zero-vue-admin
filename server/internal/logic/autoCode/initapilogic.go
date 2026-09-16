package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitAPILogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitAPILogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitAPILogic {
	return &InitAPILogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// InitAPI 初始化 API
// 根据代码生成历史记录初始化 API 数据
// TODO: 实现完整的 API 初始化逻辑
func (l *InitAPILogic) InitAPI(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Msg:  fmt.Sprintf("API 初始化成功"),
	}, nil
}
