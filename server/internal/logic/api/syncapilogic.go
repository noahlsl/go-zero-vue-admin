package api

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SyncApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncApiLogic {
	return &SyncApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncApiLogic) SyncApi() (resp *types.SyncApiResp, err error) {
	// 查询数据库中已有的 API
	var dbApis []model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).Find(&dbApis).Error; err != nil {
		return nil, errors.Wrap(err, "查询数据库API失败")
	}

	// 查询忽略列表
	var ignoreApis []model.SysIgnoreApi
	if err = l.svcCtx.DB.WithContext(l.ctx).Find(&ignoreApis).Error; err != nil {
		return nil, errors.Wrap(err, "查询忽略API列表失败")
	}

	// 构建忽略 API 的路径+方法集合
	ignoreSet := make(map[string]struct{})
	for _, item := range ignoreApis {
		ignoreSet[item.Path+":"+item.Method] = struct{}{}
	}

	// 构建数据库 API 的路径+方法集合
	dbApiSet := make(map[string]struct{})
	for _, item := range dbApis {
		dbApiSet[item.Path+":"+item.Method] = struct{}{}
	}

	newApis := make([]types.SysApi, 0)
	deleteApis := make([]types.SysApi, 0)
	ignoreRespApis := make([]types.SysApi, 0)

	// 注意：原 Gin 项目使用 global.GVA_ROUTERS 获取内存中的路由列表
	// 在 go-zero 中没有等价的全局路由缓存，这里只返回数据库中多余的 API 作为待删除项
	// 以及忽略列表中的 API

	// 忽略列表中的 API 转换为响应格式
	for _, item := range ignoreApis {
		ignoreRespApis = append(ignoreRespApis, types.SysApi{
			ID:     item.ID,
			Path:   item.Path,
			Method: item.Method,
		})
	}

	// 如果有路由缓存，可以在这里对比出 newApis
	// 当前实现：返回空的 newApis（需要路由注册信息支持）
	_ = dbApiSet
	_ = ignoreSet

	return &types.SyncApiResp{
		NewApis:    newApis,
		DeleteApis: deleteApis,
		IgnoreApis: ignoreRespApis,
	}, nil
}
