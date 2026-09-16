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

type PreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewLogic {
	return &PreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Preview 预览代码生成结果
// 根据表名查询列信息，生成预览代码片段
func (l *PreviewLogic) Preview(req *types.PreviewReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.TableName) == "" {
		return nil, errors.New("表名不能为空")
	}

	// 查询当前数据库中的列信息
	columns, err := l.getColumnsForPreview(req.TableName)
	if err != nil {
		return nil, errors.Wrapf(err, "预览表 %s 失败", req.TableName)
	}

	// 生成预览数据
	preview := map[string]interface{}{
		"tableName":  req.TableName,
		"columns":    columns,
		"structName": toFieldName(req.TableName),
	}

	return &types.Response{
		Code: 0,
		Data: preview,
		Msg:  fmt.Sprintf("表 %s 预览成功，共 %d 个字段", req.TableName, len(columns)),
	}, nil
}

func (l *PreviewLogic) getColumnsForPreview(tableName string) ([]types.ColumnInfo, error) {
	logic := NewGetColumnLogic(l.ctx, l.svcCtx)
	return logic.GetColumnByTable(tableName, "")
}
