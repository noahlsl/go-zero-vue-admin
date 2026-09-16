package sysVersion

import (
	"context"
	"fmt"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ImportVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportVersionLogic {
	return &ImportVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportVersionLogic) ImportVersion(req *types.ImportVersionReq) (resp *types.Response, err error) {
	if req.Version == "" || req.Content == "" {
		return nil, errors.New("版本信息格式错误")
	}

	// 创建导入版本记录
	now := time.Now().Format("20060102150405")
	version := model.SysVersion{
		VersionName: strPtr(req.Version),
		VersionCode: strPtr(fmt.Sprintf("%s_imported_%s", req.Version, now)),
		Description: strPtr(fmt.Sprintf("导入版本: %s", req.Version)),
		VersionData: strPtr(req.Content),
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Create(&version).Error; err != nil {
		return nil, errors.Wrap(err, "保存导入版本记录失败")
	}

	logx.WithContext(l.ctx).Infow("版本导入成功",
		logx.Field("module", "sysVersion"),
		logx.Field("action", "import_version"),
		logx.Field("version", req.Version),
	)

	return &types.Response{
		Code: 0,
		Msg:  "导入成功",
	}, nil
}

func strPtr(s string) *string {
	return &s
}
