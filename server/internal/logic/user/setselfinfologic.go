package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetSelfInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSelfInfoLogic {
	return &SetSelfInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSelfInfoLogic) SetSelfInfo(req *types.ChangeUserInfoReq) (resp *types.Response, err error) {
	// 1. 从 context 获取当前登录用户 ID
	userId := middleware.GetUserId(l.ctx)
	if userId == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	// 2. 更新用户信息（只允许更新自己的信息）
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
		}).Error; err != nil {
		l.Logger.Errorw("设置自身信息失败",
			logx.Field("error", err),
			logx.Field("user_id", userId),
			logx.Field("module", "user"),
			logx.Field("action", "set_self_info"),
		)
		return nil, errors.Wrap(err, "设置自身信息失败")
	}

	return &types.Response{Msg: "设置成功"}, nil
}
