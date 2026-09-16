package casbin

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyPathByAuthorityIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyPathByAuthorityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyPathByAuthorityIdLogic {
	return &GetPolicyPathByAuthorityIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPolicyPathByAuthorityIdLogic) GetPolicyPathByAuthorityId(req *types.GetPolicyPathByAuthorityIdReq) (resp *types.CasbinRes, err error) {
	var casbinRules []model.SysCasbinRule
	if err = l.svcCtx.DB.WithContext(l.ctx).
		Where("ptype = ? AND v0 = ?", "p", req.AuthorityId).
		Find(&casbinRules).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色权限策略失败")
	}

	list := make([]types.CasbinInfo, 0, len(casbinRules))
	for _, rule := range casbinRules {
		list = append(list, types.CasbinInfo{
			Path:   rule.V1,
			Method: rule.V2,
		})
	}

	return &types.CasbinRes{List: list}, nil
}
