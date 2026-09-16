package file

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditFileNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditFileNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditFileNameLogic {
	return &EditFileNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditFileNameLogic) EditFileName(req *types.EditFileNameReq) (resp *types.Response, err error) {
	var file model.ExaFileUploadAndDownload
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&file).Error; err != nil {
		return nil, errors.Wrap(err, "文件记录不存在")
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.ExaFileUploadAndDownload{}).
		Where("id = ?", req.ID).Update("name", req.Name).Error; err != nil {
		logx.Errorw("编辑文件名失败",
			logx.Field("error", err.Error()),
			logx.Field("id", req.ID),
			logx.Field("module", "file"),
			logx.Field("action", "editFileName"),
		)
		return nil, errors.Wrap(err, "编辑文件名失败")
	}

	return &types.Response{Msg: "编辑成功"}, nil
}
