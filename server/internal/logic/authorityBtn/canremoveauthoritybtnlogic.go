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

type CanRemoveAuthorityBtnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCanRemoveAuthorityBtnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CanRemoveAuthorityBtnLogic {
	return &CanRemoveAuthorityBtnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CanRemoveAuthorityBtnLogic) CanRemoveAuthorityBtn(req *types.GetById) (resp *types.Response, err error) {
	var count int64
	err = l.svcCtx.DB.WithContext(l.ctx).
		Model(&model.SysAuthorityBtn{}).
		Where("sys_base_menu_btn_id = ?", req.ID).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "查询按钮权限使用情况失败")
	}
	if count > 0 {
		return nil, errors.New("此按钮正在被使用无法删除")
	}
	return &types.Response{Msg: "删除成功"}, nil
}
