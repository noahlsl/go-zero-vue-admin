package autoCode

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDBLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type dbNameRow struct {
	DatabaseName string `gorm:"column:database"`
}

func NewGetDBLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDBLogic {
	return &GetDBLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDBLogic) GetDB() (resp []types.DBInfo, err error) {
	var rows []dbNameRow
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME NOT IN ('information_schema','mysql','performance_schema','sys')"
	if err = l.svcCtx.DB.WithContext(l.ctx).Raw(sql).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "获取数据库列表失败")
	}

	resp = make([]types.DBInfo, 0, len(rows))
	for i, row := range rows {
		resp = append(resp, types.DBInfo{
			ID:   uint(i + 1),
			Name: row.DatabaseName,
		})
	}
	return resp, nil
}
