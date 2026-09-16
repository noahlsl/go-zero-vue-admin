package skills

import (
	"context"
	"strings"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveResourceLogic {
	return &SaveResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveResourceLogic) SaveResource(req *types.SaveResourceReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("资源文件名不能为空")
	}

	if req.ID != 0 {
		var res model.SysSkillResource
		if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&res).Error; err != nil {
			return nil, errors.Wrap(err, "资源文件不存在")
		}
		res.Name = req.Name
		res.Type = req.Type
		res.Content = req.Content
		if err = l.svcCtx.DB.WithContext(l.ctx).Save(&res).Error; err != nil {
			return nil, errors.Wrap(err, "更新资源文件失败")
		}
		return &types.Response{Code: 0, Msg: "更新成功"}, nil
	}

	res := model.SysSkillResource{
		Name:    req.Name,
		Type:    req.Type,
		Content: req.Content,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&res).Error; err != nil {
		return nil, errors.Wrap(err, "创建资源文件失败")
	}
	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
