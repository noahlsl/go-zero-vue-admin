package authority

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/logic/conv"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAuthorityLogic {
	return &UpdateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAuthorityLogic) UpdateAuthority(req *types.UpdateAuthorityReq) (resp *types.Response, err error) {
	// 1. 查询角色是否存在
	var oldAuthority model.SysAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", req.AuthorityId).First(&oldAuthority).Error; err != nil {
		l.Logger.Errorw("查询角色数据失败",
			logx.Field("error", err),
			logx.Field("authority_id", req.AuthorityId),
			logx.Field("module", "authority"),
			logx.Field("action", "update"),
		)
		return nil, errors.New("查询角色数据失败")
	}

	// 2. 更新角色字段
	updateData := map[string]interface{}{
		"authority_name": req.AuthorityName,
		"parent_id":      &req.ParentId,
		"default_router": req.DefaultRouter,
		"data_scope":     req.DataScope,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&oldAuthority).Updates(updateData).Error; err != nil {
		return nil, errors.Wrap(err, "更新角色失败")
	}

	// 3. 重新查询更新后的数据
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", req.AuthorityId).First(&oldAuthority).Error; err != nil {
		return nil, errors.Wrap(err, "查询更新后角色失败")
	}

	resp = &types.Response{
		Code: 0,
		Msg:  "更新成功",
		Data: conv.ToTypesAuthority(oldAuthority),
	}
	return resp, nil
}
