package menu

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseMenuLogic {
	return &UpdateBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBaseMenuLogic) UpdateBaseMenu(req *types.UpdateMenuReq) (resp *types.Response, err error) {
	db := l.svcCtx.DB

	// 检查 name 是否重复（排除自身）
	var count int64
	if err = db.Model(&model.SysBaseMenu{}).
		Where("id <> ? AND name = ?", req.ID, req.Name).
		Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "查询菜单名称失败")
	}
	if count > 0 {
		return nil, errors.New("存在相同 name，请修改 name")
	}

	// 事务更新
	err = db.Transaction(func(tx *gorm.DB) error {
		// 先删除旧的参数和按钮
		if err := tx.Unscoped().Where("sys_base_menu_id = ?", req.ID).
			Delete(&model.SysBaseMenuParameter{}).Error; err != nil {
			return errors.Wrap(err, "删除旧菜单参数失败")
		}
		if err := tx.Unscoped().Where("sys_base_menu_id = ?", req.ID).
			Delete(&model.SysBaseMenuBtn{}).Error; err != nil {
			return errors.Wrap(err, "删除旧菜单按钮失败")
		}

		// 创建新的参数
		if len(req.Button) > 0 {
			btns := make([]model.SysBaseMenuBtn, 0, len(req.Button))
			for _, b := range req.Button {
				btns = append(btns, model.SysBaseMenuBtn{
					Name:          b.Name,
					Desc:          b.Desc,
					SysBaseMenuID: req.ID,
				})
			}
			if err := tx.Create(&btns).Error; err != nil {
				return errors.Wrap(err, "创建菜单按钮失败")
			}
		}

		// 更新菜单基础字段
		updates := map[string]interface{}{
			"parent_id":         req.ParentId,
			"path":              req.Path,
			"name":              req.Name,
			"component":         req.Component,
			"sort":              req.Sort,
			"meta_keep_alive":   req.KeepAlive,
			"meta_default_menu": req.DefaultMenu,
			"meta_title":        req.Title,
			"meta_icon":         req.Icon,
			"meta_close_tab":    req.CloseTab,
			"hidden":            req.Hidden,
		}
		if err := tx.Model(&model.SysBaseMenu{}).Where("id = ?", req.ID).
			Updates(updates).Error; err != nil {
			return errors.Wrap(err, "更新菜单失败")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "更新菜单事务失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "更新成功",
	}, nil
}
