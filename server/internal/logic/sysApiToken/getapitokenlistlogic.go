package sysApiToken

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiTokenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiTokenListLogic {
	return &GetApiTokenListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiTokenListLogic) GetApiTokenList(req *types.GetApiTokenListReq) (resp *types.PageResult, err error) {
	// 校验分页参数，防止 0 值或负数导致异常查询
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 1. 计算分页参数
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysApiToken{})

	// 2. 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询API Token总数失败")
	}

	// 3. 分页查询
	var list []model.SysApiToken
	if err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&list).Error; err != nil {
		l.Logger.Errorw("查询API Token列表失败",
			logx.Field("error", err),
			logx.Field("module", "sysApiToken"),
			logx.Field("action", "get_list"),
		)
		return nil, errors.Wrap(err, "查询API Token列表失败")
	}

	// 4. 转换为响应类型
	typeList := make([]types.SysApiToken, 0, len(list))
	for _, item := range list {
		typeList = append(typeList, types.SysApiToken{
			ID:        item.ID,
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
			UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
			Token:     item.Token,
			Name:      item.Remark,
		})
	}

	return &types.PageResult{
		List:     typeList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
