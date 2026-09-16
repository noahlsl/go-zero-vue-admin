package autoCode

import (
	"context"
	"strings"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePackageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePackageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePackageLogic {
	return &CreatePackageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePackageLogic) CreatePackage(req *types.CreatePackageReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.PackageName) == "" {
		return nil, errors.New("包名不能为空")
	}

	// 检查是否已存在
	var count int64
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysAutoCodePackage{}).
		Where("package_name = ?", req.PackageName).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "检查包名失败")
	}
	if count > 0 {
		return nil, errors.New("包名已存在")
	}

	pkg := model.SysAutoCodePackage{
		PackageName: req.PackageName,
		Label:       req.PackageName,
		Template:    "package",
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&pkg).Error; err != nil {
		return nil, errors.Wrap(err, "创建包失败")
	}

	return &types.Response{
		Code: 0,
		Data: pkg,
		Msg:  "包创建成功",
	}, nil
}
