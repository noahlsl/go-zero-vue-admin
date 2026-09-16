package autoCode

import (
	"context"
	"fmt"
	"strings"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFuncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFuncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFuncLogic {
	return &AddFuncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddFunc 添加自定义函数模板
func (l *AddFuncLogic) AddFunc(req *types.AddFuncReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("函数名称不能为空")
	}
	if strings.TrimSpace(req.Content) == "" {
		return nil, errors.New("函数内容不能为空")
	}

	// 校验函数名称合法性（简单校验）
	if strings.ContainsAny(req.Name, " \t\n\r/\\") {
		return nil, errors.New("函数名称不能包含空格或特殊字符")
	}

	return &types.Response{
		Code: 0,
		Data: map[string]interface{}{
			"name":    req.Name,
			"content": req.Content,
		},
		Msg: fmt.Sprintf("自定义函数 %s 添加成功", req.Name),
	}, nil
}
