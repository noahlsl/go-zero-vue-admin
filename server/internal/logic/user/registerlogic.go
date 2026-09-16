package user

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.SysUserResponse, err error) {
	// 1. 校验用户名是否已存在
	var existUser model.SysUser
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("username = ?", req.Username).First(&existUser).Error; err == nil {
		return nil, errors.New("用户名已注册")
	} else if err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "查询用户是否存在失败")
	}

	// 2. bcrypt 加密密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "密码加密失败")
	}

	// 3. 构造角色关联
	var authorities []model.SysAuthority
	for _, authorityId := range req.AuthorityIds {
		authorities = append(authorities, model.SysAuthority{
			AuthorityId: authorityId,
		})
	}

	// 4. 创建用户
	user := model.SysUser{
		UUID:        uuid.New(),
		Username:    req.Username,
		NickName:    req.NickName,
		Password:    string(hashedPwd),
		HeaderImg:   req.HeaderImg,
		AuthorityId: req.AuthorityId,
		Authorities: authorities,
		Enable:      req.Enable,
		Phone:       req.Phone,
		Email:       req.Email,
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&user).Error; err != nil {
		l.Logger.Errorw("注册用户失败",
			logx.Field("error", err),
			logx.Field("module", "user"),
			logx.Field("action", "register"),
		)
		return nil, errors.Wrap(err, "注册用户失败")
	}

	// 5. 从数据库重新查询用户信息（含关联角色）
	if err = l.svcCtx.DB.WithContext(l.ctx).Preload("Authorities").Preload("Authority").
		First(&user, user.ID).Error; err != nil {
		return nil, errors.Wrap(err, "查询注册用户信息失败")
	}

	// 6. 构造响应
	resp = &types.SysUserResponse{
		User: conv.ToTypesUser(user),
	}

	return resp, nil
}
