package api

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetApiRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetApiRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetApiRolesLogic {
	return &SetApiRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetApiRolesLogic) SetApiRoles(req *types.SetApiRolesReq) (resp *types.Response, err error) {
	if req.Path == "" || req.Method == "" {
		return nil, errors.New("API路径和请求方法不能为空")
	}

	// 先删除该 path + method 对应的所有 Casbin 策略
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("v1 = ? AND v2 = ?", req.Path, req.Method).
		Delete(&model.SysCasbinRule{}).Error; err != nil {
		return nil, errors.Wrap(err, "清除旧策略失败")
	}

	// 批量插入新的角色关联
	if len(req.AuthorityIds) > 0 {
		rules := make([]model.SysCasbinRule, 0, len(req.AuthorityIds))
		for _, authorityId := range req.AuthorityIds {
			rules = append(rules, model.SysCasbinRule{
				Ptype: "p",
				V0:    strconv.Itoa(int(authorityId)),
				V1:    req.Path,
				V2:    req.Method,
			})
		}
		if err = l.svcCtx.DB.WithContext(l.ctx).Create(&rules).Error; err != nil {
			return nil, errors.Wrap(err, "插入新策略失败")
		}
	}

	return &types.Response{Code: 0, Msg: "设置成功"}, nil
}
