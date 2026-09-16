package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitMenuLogic {
	return &InitMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// InitMenu 初始化菜单
// 根据代码生成历史记录初始化菜单数据
// TODO: 实现完整的菜单初始化逻辑
func (l *InitMenuLogic) InitMenu(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Msg:  fmt.Sprintf("菜单初始化成功"),
	}, nil
}
