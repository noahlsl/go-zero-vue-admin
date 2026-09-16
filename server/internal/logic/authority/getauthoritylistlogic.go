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

type GetAuthorityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthorityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorityListLogic {
	return &GetAuthorityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthorityListLogic) GetAuthorityList(req *types.GetAuthorityListReq) (resp *types.PageResult, err error) {
	// 1. 分页查询角色列表
	var authorities []model.SysAuthority
	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysAuthority{})

	// 关键字搜索
	if req.Keyword != "" {
		db = db.Where("authority_name LIKE ?", "%"+req.Keyword+"%")
	}

	// 查询总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色总数失败")
	}

	// 分页查询
	page, pageSize := l.parsePageParams(req.Page, req.PageSize)
	if err = db.Offset((page - 1) * pageSize).Limit(pageSize).
		Order("authority_id ASC").Find(&authorities).Error; err != nil {
		return nil, errors.Wrap(err, "查询角色列表失败")
	}

	// 2. 构造响应列表
	list := make([]types.SysAuthority, 0, len(authorities))
	for _, a := range authorities {
		list = append(list, conv.ToTypesAuthority(a))
	}

	resp = &types.PageResult{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	return resp, nil
}

func (l *GetAuthorityListLogic) parsePageParams(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
