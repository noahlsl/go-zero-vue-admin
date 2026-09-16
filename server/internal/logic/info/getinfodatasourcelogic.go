package info

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoDataSourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoDataSourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoDataSourceLogic {
	return &GetInfoDataSourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetInfoDataSource 获取公告关联的用户数据源列表
func (l *GetInfoDataSourceLogic) GetInfoDataSource() (resp []types.InfoDataSourceItem, err error) {
	var users []model.SysUser
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Select("id, nick_name").
		Find(&users).Error; err != nil {
		return nil, errors.Wrap(err, "查询用户列表失败")
	}

	resp = make([]types.InfoDataSourceItem, 0, len(users))
	for _, u := range users {
		resp = append(resp, types.InfoDataSourceItem{
			Label: u.NickName,
			Value: u.ID,
		})
	}

	return resp, nil
}
