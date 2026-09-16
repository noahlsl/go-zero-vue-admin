package file

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileListLogic {
	return &GetFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileListLogic) GetFileList(req *types.GetExaFileListReq) (resp *types.PageResult, err error) {
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

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.ExaFileUploadAndDownload{}).Order("id desc")

	// 2. 关键字搜索
	if len(req.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}

	// 3. 统计总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "查询文件总数失败")
	}

	// 4. 分页查询
	var list []model.ExaFileUploadAndDownload
	if err = db.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		l.Logger.Errorw("查询文件列表失败",
			logx.Field("error", err),
			logx.Field("module", "file"),
			logx.Field("action", "get_list"),
		)
		return nil, errors.Wrap(err, "查询文件列表失败")
	}

	// 5. 转换为响应类型
	typeList := make([]types.ExaFileUploadAndDownload, 0, len(list))
	for _, item := range list {
		typeList = append(typeList, types.ExaFileUploadAndDownload{
			ID:        item.ID,
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
			UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
			Name:      item.Name,
			Url:       item.Url,
		})
	}

	return &types.PageResult{
		List:     typeList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
