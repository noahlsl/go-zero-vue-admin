package sysVersion

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FindSysVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSysVersionLogic {
	return &FindSysVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSysVersionLogic) FindSysVersion(req *types.GetById) (resp *types.SysVersionRes, err error) {
	var version model.SysVersion
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("id = ?", req.ID).
		First(&version).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("版本记录不存在")
		}
		return nil, errors.Wrap(err, "查询版本记录失败")
	}

	resp = &types.SysVersionRes{
		Version: types.SysVersion{
			ID:          version.ID,
			CreatedAt:   version.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   version.UpdatedAt.Format(time.RFC3339),
			VersionName: version.VersionName,
			VersionCode: version.VersionCode,
			Description: version.Description,
			VersionData: version.VersionData,
		},
	}
	return resp, nil
}

func derefStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
