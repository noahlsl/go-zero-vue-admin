package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoRes, err error) {
	// 1. 从 JWT context 提取用户 ID
	userId := middleware.GetUserId(l.ctx)
	if userId == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	// 2. 查询用户信息，关联角色
	var user model.SysUser
	if err = l.svcCtx.DB.WithContext(l.ctx).Preload("Authorities").Preload("Authority").
		First(&user, userId).Error; err != nil {
		l.Logger.Errorw("查询用户信息失败",
			logx.Field("error", err),
			logx.Field("user_id", userId),
			logx.Field("module", "user"),
			logx.Field("action", "get_user_info"),
		)
		return nil, errors.Wrap(err, "查询用户信息失败")
	}

	// 3. 校验角色默认路由，无效时回退 404，避免前端跳转首页失败
	conv.EnsureDefaultRouter(l.ctx, l.svcCtx.DB, &user.Authority)

	// 4. 构造响应，字段名与原 Gin 后端 gin.H{"userInfo": ...} 保持一致
	resp = &types.UserInfoRes{
		UserInfo: conv.ToTypesUser(user),
	}

	return resp, nil
}
