package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.GetById) (resp *types.Response, err error) {
	// 1. 从 context 获取当前登录用户 ID，禁止删除自己
	userId := middleware.GetUserId(l.ctx)
	if userId == uint(req.ID) {
		return nil, errors.New("删除失败, 无法删除自己")
	}

	// 2. 事务执行软删除
	if err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 软删除用户
		if err := tx.Where("id = ?", req.ID).Delete(&model.SysUser{}).Error; err != nil {
			return errors.Wrap(err, "删除用户失败")
		}
		// 删除用户角色关联
		if err := tx.Where("sys_user_id = ?", req.ID).
			Delete(&model.SysUserAuthority{}).Error; err != nil {
			return errors.Wrap(err, "删除用户角色关联失败")
		}
		return nil
	}); err != nil {
		l.Logger.Errorw("删除用户失败",
			logx.Field("error", err),
			logx.Field("user_id", req.ID),
			logx.Field("operator_id", userId),
			logx.Field("module", "user"),
			logx.Field("action", "delete_user"),
		)
		return nil, err
	}

	return &types.Response{Msg: "删除成功"}, nil
}
