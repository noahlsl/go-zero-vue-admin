package autoCode

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplatesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplatesLogic {
	return &GetTemplatesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetTemplates 获取代码生成模板列表
// TODO: 实现从文件系统或数据库读取模板列表
func (l *GetTemplatesLogic) GetTemplates() (resp *types.Response, err error) {
	// 返回内置模板列表
	templates := []map[string]interface{}{
		{"name": "package", "desc": "分层模板 (api/service/router)"},
		{"name": "page", "desc": "表单生成器模板"},
		{"name": "plugin", "desc": "插件模板"},
	}

	return &types.Response{
		Code: 0,
		Data: templates,
		Msg:  fmt.Sprintf("获取模板列表成功，共 %d 个模板", len(templates)),
	}, nil
}
