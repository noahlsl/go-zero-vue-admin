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

type CreateResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceLogic {
	return &CreateResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateResourceLogic) CreateResource(req *types.CreateResourceReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("资源文件名不能为空")
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
