package sysVersion

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DownloadVersionJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadVersionJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadVersionJsonLogic {
	return &DownloadVersionJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadVersionJsonLogic) DownloadVersionJson() (resp *types.Response, err error) {
	// 获取最新的版本记录
	var version model.SysVersion
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Order("id DESC").
		First(&version).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("没有可用的版本记录")
		}
		return nil, errors.Wrap(err, "获取版本记录失败")
	}

	jsonData := derefStr(version.VersionData)
	if jsonData == "" {
		jsonData = "{}"
	}

	logx.WithContext(l.ctx).Infow("版本JSON下载成功",
		logx.Field("module", "sysVersion"),
		logx.Field("action", "download_version_json"),
		logx.Field("id", version.ID),
	)

	return &types.Response{
		Code: 0,
		Data: jsonData,
		Msg:  "下载成功",
	}, nil
}
