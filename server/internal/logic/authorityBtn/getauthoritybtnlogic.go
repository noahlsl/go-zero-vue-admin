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
)

type GetAuthorityBtnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthorityBtnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorityBtnLogic {
	return &GetAuthorityBtnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthorityBtnLogic) GetAuthorityBtn(req *types.GetAuthorityBtnReq) (resp *types.AuthorityBtnRes, err error) {
	var authorityBtns []model.SysAuthorityBtn
	err = l.svcCtx.DB.WithContext(l.ctx).
		Where("authority_id = ? AND sys_menu_id = ?", req.AuthorityId, req.MenuId).
		Find(&authorityBtns).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询按钮权限失败")
	}

	selected := make([]uint, 0, len(authorityBtns))
	for _, btn := range authorityBtns {
		selected = append(selected, btn.SysBaseMenuBtnID)
	}

	return &types.AuthorityBtnRes{
		Selected: selected,
	}, nil
}
