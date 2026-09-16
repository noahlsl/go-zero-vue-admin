// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package authorityBtn

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SetAuthorityBtnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAuthorityBtnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAuthorityBtnLogic {
	return &SetAuthorityBtnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAuthorityBtnLogic) SetAuthorityBtn(req *types.SetAuthorityBtnReq) (resp *types.Response, err error) {
	err = l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除该角色+菜单下已有的按钮权限
		if err = tx.Where("authority_id = ? AND sys_menu_id = ?", req.AuthorityId, req.MenuId).
			Delete(&model.SysAuthorityBtn{}).Error; err != nil {
			return errors.Wrap(err, "删除旧按钮权限失败")
		}

		// 批量创建新的按钮权限
		if len(req.Selected) > 0 {
			authorityBtns := make([]model.SysAuthorityBtn, 0, len(req.Selected))
			for _, btnId := range req.Selected {
				authorityBtns = append(authorityBtns, model.SysAuthorityBtn{
					AuthorityId:      req.AuthorityId,
					SysMenuID:        req.MenuId,
					SysBaseMenuBtnID: btnId,
				})
			}
			if err = tx.Create(&authorityBtns).Error; err != nil {
				return errors.Wrap(err, "创建按钮权限失败")
			}
		}
		return nil
	})
	if err != nil {
		logx.Errorw("设置按钮权限失败",
			logx.Field("error", err.Error()),
			logx.Field("authorityId", req.AuthorityId),
			logx.Field("menuId", req.MenuId),
			logx.Field("module", "authorityBtn"),
			logx.Field("action", "setAuthorityBtn"),
		)
		return nil, err
	}
	return &types.Response{Msg: "分配成功"}, nil
}
