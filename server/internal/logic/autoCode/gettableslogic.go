package autoCode

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTablesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type tableNameRow struct {
	TableName string `gorm:"column:table_name"`
}

func NewGetTablesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTablesLogic {
	return &GetTablesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTablesLogic) GetTables() (resp []types.TableInfo, err error) {
	// 使用当前连接的数据库获取表列表
	var rows []tableNameRow
	sql := "SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_type = 'BASE TABLE'"
	if err = l.svcCtx.DB.WithContext(l.ctx).Raw(sql).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "获取表列表失败")
	}

	resp = make([]types.TableInfo, 0, len(rows))
	for i, row := range rows {
		resp = append(resp, types.TableInfo{
			ID:        uint(i + 1),
			TableName: row.TableName,
		})
	}
	return resp, nil
}
