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

type CreateReferenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReferenceLogic {
	return &CreateReferenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReferenceLogic) CreateReference(req *types.CreateReferenceReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("参考文件名不能为空")
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
