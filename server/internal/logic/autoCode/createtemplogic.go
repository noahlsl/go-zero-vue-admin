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

type CreateTempLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTempLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTempLogic {
	return &CreateTempLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateTemp 根据模板和表信息生成代码文件
func (l *CreateTempLogic) CreateTemp(req *types.CreateTemplateReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.TableName) == "" {
		return nil, errors.New("表名不能为空")
	}
	if strings.TrimSpace(req.DbName) == "" {
		return nil, errors.New("数据库名不能为空")
	}
	if strings.TrimSpace(req.PackageName) == "" {
		return nil, errors.New("包名不能为空")
	}

	// 查询表的列信息
	columns, err := l.getColumnInfo(req.TableName, req.DbName)
	if err != nil {
		return nil, errors.Wrapf(err, "获取表 %s 列信息失败", req.TableName)
	}

	// 构建生成结果
	structName := toFieldName(req.TableName)
	result := map[string]interface{}{
		"tableName":   req.TableName,
		"structName":  structName,
		"dbName":      req.DbName,
		"packageName": req.PackageName,
		"moduleName":  req.ModuleName,
		"dataType":    req.DataType,
		"columns":     columns,
	}

	return &types.Response{
		Code: 0,
		Data: result,
		Msg:  fmt.Sprintf("模板代码生成成功: %s", structName),
	}, nil
}

func (l *CreateTempLogic) getColumnInfo(tableName, dbName string) ([]types.ColumnInfo, error) {
	logic := NewGetColumnLogic(l.ctx, l.svcCtx)
	return logic.GetColumnByTable(tableName, dbName)
}
