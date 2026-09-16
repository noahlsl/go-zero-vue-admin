package authority

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAuthorityLogic {
	return &CreateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAuthorityLogic) CreateAuthority(req *types.CreateAuthorityReq) (resp *types.Response, err error) {
	// 1. 校验角色ID是否已存在
	var existAuth model.SysAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", req.AuthorityId).First(&existAuth).Error; err == nil {
		return nil, errors.New("存在相同角色id")
	} else if err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "查询角色是否已存在失败")
	}

	// 2. 创建角色并关联默认菜单
	auth := model.SysAuthority{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      &req.ParentId,
		DefaultRouter: req.DefaultRouter,
		DataScope:     req.DataScope,
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&auth).Error; e != nil {
			return errors.Wrap(e, "创建角色失败")
		}

		// 关联默认菜单（仪表盘，菜单ID=1）
		if e := tx.Exec(
			"INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) VALUES (?, ?)",
			req.AuthorityId, 1,
		).Error; e != nil {
			return errors.Wrap(e, "关联默认菜单失败")
		}

		return nil
	})
	if err != nil {
		l.Logger.Errorw("创建角色失败",
			logx.Field("error", err),
			logx.Field("authority_id", req.AuthorityId),
			logx.Field("module", "authority"),
			logx.Field("action", "create"),
		)
		return nil, err
	}

	resp = &types.Response{
		Code: 0,
		Msg:  "创建成功",
		Data: conv.ToTypesAuthority(auth),
	}
	return resp, nil
}
