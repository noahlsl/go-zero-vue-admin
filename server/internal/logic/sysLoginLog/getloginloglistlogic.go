package sysLoginLog

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogListLogic {
	return &GetLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogListLogic) GetLoginLogList(req *types.GetLoginLogListReq) (resp *types.PageResult, err error) {
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

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysLoginLog{})

	// 2. 条件筛选
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+req.Ip+"%")
	}
	if req.Status {
		db = db.Where("status = ?", req.Status)
	}

	// 3. 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询登录日志总数失败")
	}

	// 4. 分页查询
	var list []model.SysLoginLog
	if err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error; err != nil {
		l.Logger.Errorw("查询登录日志列表失败",
			logx.Field("error", err),
			logx.Field("module", "sysLoginLog"),
			logx.Field("action", "get_list"),
		)
		return nil, errors.Wrap(err, "查询登录日志列表失败")
	}

	// 5. 转换为响应类型
	typeList := make([]types.SysLoginLog, 0, len(list))
	for _, item := range list {
		typeList = append(typeList, modelLoginLogToTypes(item))
	}

	return &types.PageResult{
		List:     typeList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
