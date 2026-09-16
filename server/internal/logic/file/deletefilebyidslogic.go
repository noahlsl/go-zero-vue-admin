package file

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileByIdsLogic {
	return &DeleteFileByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileByIdsLogic) DeleteFileByIds(req *types.IdsReq) (resp *types.Response, err error) {
	// 1. 参数校验
	if len(req.Ids) == 0 {
		return nil, errors.New("删除列表不能为空")
	}

	// 2. 批量物理删除
	if err = l.svcCtx.DB.WithContext(l.ctx).Unscoped().Where("id in ?", req.Ids).Delete(&model.ExaFileUploadAndDownload{}).Error; err != nil {
		l.Logger.Errorw("批量删除文件记录失败",
			logx.Field("error", err),
			logx.Field("ids", req.Ids),
			logx.Field("module", "file"),
			logx.Field("action", "delete_by_ids"),
		)
		return nil, errors.Wrap(err, "批量删除文件记录失败")
	}

	return &types.Response{Msg: "删除成功"}, nil
}
