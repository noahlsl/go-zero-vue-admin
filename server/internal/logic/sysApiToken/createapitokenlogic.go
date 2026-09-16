package sysApiToken

import (
	"context"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiTokenLogic {
	return &CreateApiTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiTokenLogic) CreateApiToken(req *types.CreateApiTokenReq) (resp *types.CreateApiTokenRes, err error) {
	// 1. 参数校验
	if req.Name == "" {
		return nil, errors.New("Token名称不能为空")
	}

	// 2. 生成唯一 Token
	tokenStr := uuid.New().String()

	// 3. 创建 API Token 记录（默认有效期30天）
	apiToken := model.SysApiToken{
		Token:     tokenStr,
		Status:    true,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		Remark:    req.Name,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&apiToken).Error; err != nil {
		l.Logger.Errorw("创建API Token失败",
			logx.Field("error", err),
			logx.Field("name", req.Name),
			logx.Field("module", "sysApiToken"),
			logx.Field("action", "create"),
		)
		return nil, errors.Wrap(err, "创建API Token失败")
	}

	return &types.CreateApiTokenRes{
		Token: types.SysApiToken{
			ID:        apiToken.ID,
			CreatedAt: apiToken.CreatedAt.Format(time.RFC3339),
			UpdatedAt: apiToken.UpdatedAt.Format(time.RFC3339),
			Token:     apiToken.Token,
			Name:      apiToken.Remark,
		},
	}, nil
}
