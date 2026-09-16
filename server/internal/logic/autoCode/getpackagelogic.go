package autoCode

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPackageLogic {
	return &GetPackageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackageLogic) GetPackage() (resp []types.PackageInfo, err error) {
	var packages []model.SysAutoCodePackage
	if err = l.svcCtx.DB.WithContext(l.ctx).Find(&packages).Error; err != nil {
		return nil, errors.Wrap(err, "获取包列表失败")
	}

	resp = make([]types.PackageInfo, 0, len(packages))
	for _, pkg := range packages {
		resp = append(resp, types.PackageInfo{
			ID:          pkg.ID,
			PackageName: pkg.PackageName,
			WebPath:     pkg.Label,
			GoModFile:   pkg.Template,
		})
	}
	return resp, nil
}
