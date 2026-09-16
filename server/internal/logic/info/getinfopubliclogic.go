package info

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoPublicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoPublicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoPublicLogic {
	return &GetInfoPublicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetInfoPublic 获取公开的公告信息
// 返回固定的成功响应，与原 Gin 后端行为保持一致
func (l *GetInfoPublicLogic) GetInfoPublic() (resp *types.GetInfoPublicRes, err error) {
	return &types.GetInfoPublicRes{
		Success: true,
		Data: types.InfoDataSourceItem{
			Label: "公开公告",
			Value: 0,
		},
	}, nil
}
