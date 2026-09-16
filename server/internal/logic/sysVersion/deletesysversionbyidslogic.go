package sysVersion

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysVersionByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysVersionByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysVersionByIdsLogic {
	return &DeleteSysVersionByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysVersionByIdsLogic) DeleteSysVersionByIds(req *types.IdsReq) (resp *types.Response, err error) {
	if len(req.Ids) == 0 {
		return nil, errors.New("ID列表不能为空")
	}

	result := l.svcCtx.DB.WithContext(l.ctx).
		Where("id IN ?", req.Ids).
		Delete(&model.SysVersion{})
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "批量删除版本记录失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
