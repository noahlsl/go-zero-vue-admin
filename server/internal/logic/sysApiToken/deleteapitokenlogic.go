package sysApiToken

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiTokenLogic {
	return &DeleteApiTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiTokenLogic) DeleteApiToken(req *types.DeleteApiTokenReq) (resp *types.Response, err error) {
	// 1. 查询 Token 记录
	var apiToken model.SysApiToken
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&apiToken).Error; err != nil {
		return nil, errors.Wrap(err, "API Token记录不存在")
	}

	// 2. 将 Token 加入 JWT 黑名单
	blacklist := model.SysJwtBlacklist{Jwt: apiToken.Token}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&blacklist).Error; err != nil {
		l.Logger.Errorw("Token加入黑名单失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysApiToken"),
			logx.Field("action", "add_to_blacklist"),
		)
		return nil, errors.Wrap(err, "Token加入黑名单失败")
	}

	// 3. 将 Token 状态设为无效
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysApiToken{}).Where("id = ?", req.ID).
		Update("status", false).Error; err != nil {
		l.Logger.Errorw("更新Token状态失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "sysApiToken"),
			logx.Field("action", "update_status"),
		)
		return nil, errors.Wrap(err, "更新Token状态失败")
	}

	return &types.Response{Msg: "删除成功"}, nil
}
