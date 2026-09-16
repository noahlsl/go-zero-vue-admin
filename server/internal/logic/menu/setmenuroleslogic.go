package menu

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SetMenuRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMenuRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenuRolesLogic {
	return &SetMenuRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMenuRolesLogic) SetMenuRoles(req *types.SetMenuRolesReq) (resp *types.Response, err error) {
	if req.MenuId == 0 {
		return nil, errors.New("菜单ID不能为空")
	}

	menuIdStr := strconv.FormatUint(uint64(req.MenuId), 10)

	// 事务：先删除所有旧关联，再插入新关联
	err = l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		// 删除该菜单所有已有的角色关联
		if err := tx.Where("sys_base_menu_id = ?", menuIdStr).
			Delete(&model.SysAuthorityMenu{}).Error; err != nil {
			return errors.Wrap(err, "删除旧角色关联失败")
		}

		// 批量插入新的关联记录
		if len(req.AuthorityIds) > 0 {
			records := make([]model.SysAuthorityMenu, 0, len(req.AuthorityIds))
			for _, authId := range req.AuthorityIds {
				records = append(records, model.SysAuthorityMenu{
					MenuId:      menuIdStr,
					AuthorityId: strconv.FormatUint(uint64(authId), 10),
				})
			}
			if err := tx.Create(&records).Error; err != nil {
				return errors.Wrap(err, "创建角色关联失败")
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "设置菜单角色事务失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "设置成功",
	}, nil
}
