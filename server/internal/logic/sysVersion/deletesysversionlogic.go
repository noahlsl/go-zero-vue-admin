package sysVersion

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysVersionLogic {
	return &DeleteSysVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysVersionLogic) DeleteSysVersion(req *types.DeleteSysVersionReq) (resp *types.Response, err error) {
	result := l.svcCtx.DB.WithContext(l.ctx).
		Delete(&model.SysVersion{}, "id = ?", req.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "删除版本记录失败")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("版本记录不存在")
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
