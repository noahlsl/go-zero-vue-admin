package sysDictionaryDetail

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysDictionaryDetailLogic {
	return &CreateSysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysDictionaryDetailLogic) CreateSysDictionaryDetail(req *types.CreateSysDictionaryDetailReq) (resp *types.Response, err error) {
	detail := model.SysDictionaryDetail{
		Label:           req.Label,
		Value:           req.Value,
		Extend:          req.Extend,
		Status:          req.Status,
		Sort:            req.Sort,
		SysDictionaryID: req.SysDictionaryID,
		ParentID:        req.ParentID,
	}

	// 根据 ParentID 计算层级和路径
	if req.ParentID != nil {
		var parent model.SysDictionaryDetail
		if err = l.svcCtx.DB.WithContext(l.ctx).First(&parent, *req.ParentID).Error; err != nil {
			return nil, errors.Wrap(err, "查询父级字典详情失败")
		}
		detail.Level = parent.Level + 1
		if parent.Path == "" {
			detail.Path = strconv.Itoa(int(parent.ID))
		} else {
			detail.Path = parent.Path + "," + strconv.Itoa(int(parent.ID))
		}
	} else {
		detail.Level = 0
		detail.Path = ""
	}

	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&detail).Error; err != nil {
		return nil, errors.Wrap(err, "创建字典详情失败")
	}

	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
