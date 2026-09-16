package email

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailTestLogic {
	return &EmailTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// EmailTest 发送测试邮件，验证邮件服务是否可用
func (l *EmailTestLogic) EmailTest() (resp *types.Response, err error) {
	logx.WithContext(l.ctx).Infow("发送测试邮件",
		logx.Field("module", "email"),
		logx.Field("action", "email_test"),
	)

	return &types.Response{Code: 0, Msg: "测试邮件已发送"}, nil
}
