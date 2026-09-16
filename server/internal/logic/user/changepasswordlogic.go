package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/middleware"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.Response, err error) {
	// 1. 从 context 获取当前登录用户 ID
	userId := middleware.GetUserId(l.ctx)
	if userId == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	// 2. 查询用户密码
	var user model.SysUser
	if err = l.svcCtx.DB.WithContext(l.ctx).Select("id, password").Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "查询用户信息失败")
	}

	// 3. 校验旧密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("原密码错误")
	}

	// 4. bcrypt 加密新密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "新密码加密失败")
	}

	// 5. 更新密码
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysUser{}).Where("id = ?", userId).
		Update("password", string(hashedPwd)).Error; err != nil {
		l.Logger.Errorw("修改密码失败",
			logx.Field("error", err),
			logx.Field("user_id", userId),
			logx.Field("module", "user"),
			logx.Field("action", "change_password"),
		)
		return nil, errors.Wrap(err, "修改密码失败")
	}

	return &types.Response{Msg: "修改成功"}, nil
}
