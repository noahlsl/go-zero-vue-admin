package menu

import (
	"context"
	"fmt"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuAuthorityLogic {
	return &AddMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuAuthorityLogic) AddMenuAuthority(req *types.AddMenuAuthorityReq) (resp *types.Response, err error) {
	db := l.svcCtx.DB

	// 校验菜单是否存在
	var menuCount int64
	if err = db.Model(&model.SysBaseMenu{}).Where("id = ?", req.MenuId).
		Count(&menuCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单失败")
	}
	if menuCount == 0 {
		return nil, errors.New(fmt.Sprintf("菜单 %d 不存在", req.MenuId))
	}

	// 校验角色是否存在
	var authCount int64
	if err = db.Model(&model.SysAuthority{}).Where("authority_id = ?", req.AuthorityId).
		Count(&authCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色失败")
	}
	if authCount == 0 {
		return nil, errors.New(fmt.Sprintf("角色 %d 不存在", req.AuthorityId))
	}

	// 检查是否已存在关联
	var existCount int64
	if err = db.Model(&model.SysAuthorityMenu{}).
		Where("sys_base_menu_id = ? AND sys_authority_authority_id = ?",
			strconv.FormatUint(uint64(req.MenuId), 10),
			strconv.FormatUint(uint64(req.AuthorityId), 10)).
		Count(&existCount).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单角色关联失败")
	}
	if existCount > 0 {
		return &types.Response{
			Code: 0,
			Msg:  "关联关系已存在",
		}, nil
	}

	// 创建关联
	record := model.SysAuthorityMenu{
		MenuId:      strconv.FormatUint(uint64(req.MenuId), 10),
		AuthorityId: strconv.FormatUint(uint64(req.AuthorityId), 10),
	}
	if err = db.Create(&record).Error; err != nil {
		return nil, errors.Wrap(err, "创建菜单角色关联失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "添加成功",
	}, nil
}
