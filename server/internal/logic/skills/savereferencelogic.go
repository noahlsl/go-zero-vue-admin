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

type SaveReferenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveReferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveReferenceLogic {
	return &SaveReferenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveReferenceLogic) SaveReference(req *types.SaveReferenceReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("参考文件名不能为空")
	}

	if req.ID != 0 {
		var ref model.SysSkillReference
		if err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.ID).First(&ref).Error; err != nil {
			return nil, errors.Wrap(err, "参考资料不存在")
		}
		ref.Name = req.Name
		ref.Content = req.Content
		if err = l.svcCtx.DB.WithContext(l.ctx).Save(&ref).Error; err != nil {
			return nil, errors.Wrap(err, "更新参考资料失败")
		}
		return &types.Response{Code: 0, Msg: "更新成功"}, nil
	}

	ref := model.SysSkillReference{
		Name:    req.Name,
		Content: req.Content,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&ref).Error; err != nil {
		return nil, errors.Wrap(err, "创建参考资料失败")
	}
	return &types.Response{Code: 0, Msg: "创建成功"}, nil
}
