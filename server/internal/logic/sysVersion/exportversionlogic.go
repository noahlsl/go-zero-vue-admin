package sysVersion

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExportVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportVersionLogic {
	return &ExportVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportVersionLogic) ExportVersion(req *types.ExportVersionReq) (resp *types.Response, err error) {
	// 获取版本记录
	var version model.SysVersion
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("id = ?", req.ID).
		First(&version).Error; err != nil {
		return nil, errors.Wrap(err, "获取版本记录失败")
	}

	// 构建导出数据
	exportData := make(map[string]interface{})
	if version.VersionData != nil && *version.VersionData != "" {
		exportData["versionData"] = *version.VersionData
	} else {
		exportData["versionData"] = "{}"
	}
	exportData["versionName"] = derefStr(version.VersionName)
	exportData["versionCode"] = derefStr(version.VersionCode)
	exportData["description"] = derefStr(version.Description)

	logx.WithContext(l.ctx).Infow("版本导出成功",
		logx.Field("module", "sysVersion"),
		logx.Field("action", "export_version"),
		logx.Field("id", req.ID),
	)

	return &types.Response{
		Code: 0,
		Data: exportData,
		Msg:  "导出成功",
	}, nil
}
