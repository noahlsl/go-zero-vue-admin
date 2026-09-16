package api

import (
	"context"
	"strings"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiGroupsLogic {
	return &GetApiGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiGroupsLogic) GetApiGroups() (resp *types.ApiGroups, err error) {
	var apis []model.SysApi
	if err = l.svcCtx.DB.WithContext(l.ctx).Find(&apis).Error; err != nil {
		return nil, errors.Wrap(err, "查询所有API失败")
	}

	groupSet := make(map[string]struct{})
	groups := make([]string, 0)
	for _, item := range apis {
		if _, ok := groupSet[item.ApiGroup]; !ok {
			groupSet[item.ApiGroup] = struct{}{}
			groups = append(groups, item.ApiGroup)
		}
	}

	return &types.ApiGroups{Groups: groups}, nil
}

// apiGroupMap 返回路径前缀到分组名的映射
func apiGroupMap(apis []model.SysApi) map[string]string {
	m := make(map[string]string)
	for _, item := range apis {
		parts := strings.Split(item.Path, "/")
		if len(parts) > 1 {
			m[parts[1]] = item.ApiGroup
		}
	}
	return m
}
