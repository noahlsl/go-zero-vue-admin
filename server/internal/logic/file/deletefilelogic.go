package file

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.GetById) (resp *types.Response, err error) {
	// 1. 校验文件记录是否存在
	var file model.ExaFileUploadAndDownload
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&file).Error; err != nil {
		return nil, errors.Wrap(err, "文件记录不存在")
	}

	// 2. 物理删除记录（Unscoped 硬删除）
	if err = l.svcCtx.DB.WithContext(l.ctx).Unscoped().Where("id = ?", req.ID).Delete(&model.ExaFileUploadAndDownload{}).Error; err != nil {
		l.Logger.Errorw("删除文件记录失败",
			logx.Field("error", err),
			logx.Field("id", req.ID),
			logx.Field("module", "file"),
			logx.Field("action", "delete"),
		)
		return nil, errors.Wrap(err, "删除文件记录失败")
	}

	return &types.Response{Msg: "删除成功"}, nil
}
