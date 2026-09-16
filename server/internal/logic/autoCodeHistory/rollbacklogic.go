package autoCodeHistory

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackLogic {
	return &RollbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Rollback 回滚代码生成
// TODO: 实现完整的回滚逻辑（删除生成的文件、回滚注入代码、更新标记）
func (l *RollbackLogic) Rollback(req *types.RollBackReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	var history model.SysAutoCodeHistory
	if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&history).Error; err != nil {
		return nil, errors.Wrap(err, "查询历史记录失败")
	}

	// 检查是否已回滚
	if history.Flag == 1 {
		return nil, errors.New("该记录已回滚")
	}

	// 更新标记为已回滚
	if err = l.svcCtx.DB.WithContext(l.ctx).Model(&model.SysAutoCodeHistory{}).
		Where("id = ?", req.ID).
		Update("flag", 1).Error; err != nil {
		return nil, errors.Wrap(err, "更新回滚标记失败")
	}

	return &types.Response{
		Code: 0,
		Msg:  "回滚成功",
	}, nil
}
