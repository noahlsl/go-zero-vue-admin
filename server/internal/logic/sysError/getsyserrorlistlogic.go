package sysError

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysErrorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysErrorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysErrorListLogic {
	return &GetSysErrorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysErrorListLogic) GetSysErrorList(req *types.GetSysErrorListReq) (resp *types.PageResult, err error) {
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

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysError{}).Order("created_at desc")

	// 2. 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询错误日志总数失败")
	}

	// 3. 分页查询
	var list []model.SysError
	if err = db.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		l.Logger.Errorw("查询错误日志列表失败",
			logx.Field("error", err),
			logx.Field("module", "sysError"),
			logx.Field("action", "get_list"),
		)
		return nil, errors.Wrap(err, "查询错误日志列表失败")
	}

	// 4. 转换为响应类型
	typeList := make([]types.SysError, 0, len(list))
	for _, item := range list {
		typeList = append(typeList, modelSysErrorToTypes(item))
	}

	return &types.PageResult{
		List:     typeList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
