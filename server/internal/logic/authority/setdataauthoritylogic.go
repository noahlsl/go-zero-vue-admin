package authority

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SetDataAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDataAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDataAuthorityLogic {
	return &SetDataAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetDataAuthorityLogic) SetDataAuthority(req *types.SetDataAuthorityReq) (resp *types.Response, err error) {
	// 1. 查询角色是否存在
	var authority model.SysAuthority
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("authority_id = ?", req.AuthorityId).First(&authority).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("角色不存在")
		}
		return nil, errors.Wrap(err, "查询角色失败")
	}

	// 2. 更新数据权限范围
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&authority).Update("data_scope", req.DataScope).Error; err != nil {
		return nil, errors.Wrap(err, "设置数据权限失败")
	}

	resp = &types.Response{
		Code: 0,
		Msg:  "设置成功",
	}
	return resp, nil
}
