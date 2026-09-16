package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) (resp *types.Response, err error) {
	// 1. bcrypt 加密新密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "密码加密失败")
	}

	// 2. 更新用户密码
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).Where("id = ?", req.ID).
		Update("password", string(hashedPwd)).Error; err != nil {
		l.Logger.Errorw("重置用户密码失败",
			logx.Field("error", err),
			logx.Field("user_id", req.ID),
			logx.Field("module", "user"),
			logx.Field("action", "reset_password"),
		)
		return nil, errors.Wrap(err, "重置用户密码失败")
	}

	return &types.Response{Msg: "重置成功"}, nil
}
