package api

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type EnterSyncApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnterSyncApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnterSyncApiLogic {
	return &EnterSyncApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnterSyncApiLogic) EnterSyncApi() error {
	// 查询数据库中所有 API
	var dbApis []model.SysApi
	if err := l.svcCtx.DB.WithContext(l.ctx).Find(&dbApis).Error; err != nil {
		return errors.Wrap(err, "查询数据库API失败")
	}

	// 查询忽略列表
	var ignoreApis []model.SysIgnoreApi
	if err := l.svcCtx.DB.WithContext(l.ctx).Find(&ignoreApis).Error; err != nil {
		return errors.Wrap(err, "查询忽略API列表失败")
	}

	_ = dbApis
	_ = ignoreApis

	// 注意：原 Gin 项目在此处对比内存路由与数据库记录，执行批量新增/删除
	// go-zero 没有等价的全局路由缓存机制，此接口暂为空操作
	// 如需同步功能，需在 config 中配置路由注册信息后实现对比逻辑

	return nil
}
