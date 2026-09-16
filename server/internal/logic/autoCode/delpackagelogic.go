package autoCode

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelPackageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPackageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPackageLogic {
	return &DelPackageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPackageLogic) DelPackage(req *types.DeletePackageReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("包ID不能为空")
	}

	result := l.svcCtx.DB.WithContext(l.ctx).Delete(&model.SysAutoCodePackage{}, req.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "删除包失败")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("包不存在")
	}

	return &types.Response{
		Code: 0,
		Msg:  "包删除成功",
	}, nil
}
