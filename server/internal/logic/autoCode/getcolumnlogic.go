package autoCode

import (
	"context"
	"database/sql"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetColumnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type columnRow struct {
	ColumnName    string `gorm:"column:column_name"`
	DataType      string `gorm:"column:data_type"`
	DataTypeLong  string `gorm:"column:data_type_long"`
	ColumnComment string `gorm:"column:column_comment"`
	PrimaryKey    int    `gorm:"column:primary_key"`
}

func NewGetColumnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetColumnLogic {
	return &GetColumnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetColumnLogic) GetColumn() (resp []types.ColumnInfo, err error) {
	// 注意：实际使用时应通过请求参数传入 tableName 和 dbName
	// 此处为简化实现，返回空列表
	resp = make([]types.ColumnInfo, 0)
	return resp, nil
}

// GetColumnByTable 获取指定表的列信息（内部方法）
func (l *GetColumnLogic) GetColumnByTable(tableName, dbName string) ([]types.ColumnInfo, error) {
	sqlQuery := `
	SELECT 
		c.COLUMN_NAME column_name,
		c.DATA_TYPE data_type,
		CASE c.DATA_TYPE
			WHEN 'longtext' THEN CAST(c.CHARACTER_MAXIMUM_LENGTH AS CHAR)
			WHEN 'varchar' THEN CAST(c.CHARACTER_MAXIMUM_LENGTH AS CHAR)
			WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
			WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
			WHEN 'int' THEN CAST(c.NUMERIC_PRECISION AS CHAR)
			WHEN 'bigint' THEN CAST(c.NUMERIC_PRECISION AS CHAR)
			ELSE '' 
		END AS data_type_long,
		c.COLUMN_COMMENT column_comment,
		CASE WHEN kcu.COLUMN_NAME IS NOT NULL THEN 1 ELSE 0 END AS primary_key
	FROM 
		INFORMATION_SCHEMA.COLUMNS c
	LEFT JOIN 
		INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu 
	ON 
		c.TABLE_SCHEMA = kcu.TABLE_SCHEMA 
		AND c.TABLE_NAME = kcu.TABLE_NAME 
		AND c.COLUMN_NAME = kcu.COLUMN_NAME 
		AND kcu.CONSTRAINT_NAME = 'PRIMARY'
	WHERE 
		c.TABLE_NAME = ? 
		AND c.TABLE_SCHEMA = ?
	ORDER BY 
		c.ORDINAL_POSITION;`

	var rows []columnRow
	if err := l.svcCtx.DB.WithContext(l.ctx).Raw(sqlQuery, tableName, dbName).Scan(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "获取表 %s 的列信息失败", tableName)
	}

	result := make([]types.ColumnInfo, 0, len(rows))
	for i, row := range rows {
		result = append(result, types.ColumnInfo{
			ID:            uint(i + 1),
			ColumnName:    row.ColumnName,
			DataType:      row.DataType,
			DataTypeLong:  row.DataTypeLong,
			ColumnComment: row.ColumnComment,
			Required:      row.PrimaryKey == 1,
			FieldJson:     toFieldJson(row.ColumnName),
			FieldName:     toFieldName(row.ColumnName),
		})
	}
	return result, nil
}

// GetColumnByName 获取单列信息（通过 SQL 查询）
func (l *GetColumnLogic) GetColumnByName(tableName, columnName, dbName string) (*types.ColumnInfo, error) {
	sqlQuery := `
	SELECT 
		c.COLUMN_NAME column_name,
		c.DATA_TYPE data_type,
		CASE c.DATA_TYPE
			WHEN 'longtext' THEN CAST(c.CHARACTER_MAXIMUM_LENGTH AS CHAR)
			WHEN 'varchar' THEN CAST(c.CHARACTER_MAXIMUM_LENGTH AS CHAR)
			WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
			WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
			WHEN 'int' THEN CAST(c.NUMERIC_PRECISION AS CHAR)
			WHEN 'bigint' THEN CAST(c.NUMERIC_PRECISION AS CHAR)
			ELSE '' 
		END AS data_type_long,
		c.COLUMN_COMMENT column_comment,
		CASE WHEN kcu.COLUMN_NAME IS NOT NULL THEN 1 ELSE 0 END AS primary_key
	FROM 
		INFORMATION_SCHEMA.COLUMNS c
	LEFT JOIN 
		INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu 
	ON 
		c.TABLE_SCHEMA = kcu.TABLE_SCHEMA 
		AND c.TABLE_NAME = kcu.TABLE_NAME 
		AND c.COLUMN_NAME = kcu.COLUMN_NAME 
		AND kcu.CONSTRAINT_NAME = 'PRIMARY'
	WHERE 
		c.TABLE_NAME = ? 
		AND c.COLUMN_NAME = ?
		AND c.TABLE_SCHEMA = ?`

	var row columnRow
	if err := l.svcCtx.DB.WithContext(l.ctx).Raw(sqlQuery, tableName, columnName, dbName).Scan(&row).Error; err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("列 %s.%s 不存在", tableName, columnName)
		}
		return nil, errors.Wrapf(err, "获取列 %s.%s 信息失败", tableName, columnName)
	}

	return &types.ColumnInfo{
		ColumnName:    row.ColumnName,
		DataType:      row.DataType,
		DataTypeLong:  row.DataTypeLong,
		ColumnComment: row.ColumnComment,
		Required:      row.PrimaryKey == 1,
		FieldJson:     toFieldJson(row.ColumnName),
		FieldName:     toFieldName(row.ColumnName),
	}, nil
}
