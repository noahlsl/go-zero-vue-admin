package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDictionaryLogic {
	return &InitDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// InitDictionary 初始化字典
// 根据代码生成历史记录初始化字典数据
// TODO: 实现完整的字典初始化逻辑
func (l *InitDictionaryLogic) InitDictionary(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Msg:  fmt.Sprintf("字典初始化成功"),
	}, nil
}
