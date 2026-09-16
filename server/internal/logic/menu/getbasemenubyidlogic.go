package menu

import (
	"context"
	"fmt"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseMenuByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseMenuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseMenuByIdLogic {
	return &GetBaseMenuByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBaseMenuByIdLogic) GetBaseMenuById(req *types.GetById) (resp *types.SysMenu, err error) {
	db := l.svcCtx.DB

	var baseMenu model.SysBaseMenu
	if err = db.Where("id = ?", req.ID).Preload("Parameters").Preload("MenuBtn").
		First(&baseMenu).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("查询菜单 %d 失败", req.ID))
	}

	// 转换参数
	params := make([]types.SysMenuParameter, 0, len(baseMenu.Parameters))
	for _, p := range baseMenu.Parameters {
		params = append(params, types.SysMenuParameter{
			ID:            p.ID,
			SysBaseMenuID: p.SysBaseMenuID,
			Type:          p.Type,
			Key:           p.Key,
			Value:         p.Value,
		})
	}

	// 转换按钮
	btns := make([]types.SysMenuBtn, 0, len(baseMenu.MenuBtn))
	for _, btn := range baseMenu.MenuBtn {
		btns = append(btns, types.SysMenuBtn{
			ID:       btn.ID,
			Name:     btn.Name,
			Desc:     btn.Desc,
			ButtonId: btn.ID,
		})
	}

	return &types.SysMenu{
		ID:        baseMenu.ID,
		MenuLevel: baseMenu.MenuLevel,
		ParentId:  baseMenu.ParentId,
		Path:      baseMenu.Path,
		Name:      baseMenu.Name,
		Component: baseMenu.Component,
		Sort:      baseMenu.Sort,
		Meta: types.Meta{
			KeepAlive:   baseMenu.KeepAlive,
			DefaultMenu: baseMenu.DefaultMenu,
			Title:       baseMenu.Title,
			Icon:        baseMenu.Icon,
			CloseTab:    baseMenu.CloseTab,
		},
		Hidden:     baseMenu.Hidden,
		Button:     btns,
		Parameters: params,
	}, nil
}
