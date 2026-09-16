package email

import (
	"context"
	"strings"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SendEmail 向指定收件人发送邮件
// TODO: 需要在 go-zero 配置中添加 SMTP 邮件配置（Host、Port、Secret、From、Nickname、IsSSL、IsLoginAuth），
//       配置就绪后使用 github.com/jordan-wright/email 库完成实际的 SMTP 发送。
func (l *SendEmailLogic) SendEmail(req *types.EmailSendReq) (resp *types.Response, err error) {
	if strings.TrimSpace(req.To) == "" {
		return nil, errors.New("收件人不能为空")
	}
	if strings.TrimSpace(req.Subject) == "" {
		return nil, errors.New("邮件主题不能为空")
	}
	if strings.TrimSpace(req.Body) == "" {
		return nil, errors.New("邮件正文不能为空")
	}

	logx.WithContext(l.ctx).Infow("发送邮件",
		logx.Field("module", "email"),
		logx.Field("action", "send_email"),
		logx.Field("to", req.To),
		logx.Field("subject", req.Subject),
	)

	// TODO: SMTP 配置就绪后，替换为实际发送逻辑：
	// e := email.NewEmail()
	// e.From = l.svcCtx.Config.Email.Nickname + " <" + l.svcCtx.Config.Email.From + ">"
	// e.To = strings.Split(req.To, ",")
	// e.Subject = req.Subject
	// e.HTML = []byte(req.Body)
	// addr := fmt.Sprintf("%s:%d", l.svcCtx.Config.Email.Host, l.svcCtx.Config.Email.Port)
	// if err = e.SendWithTLS(addr, smtp.PlainAuth("", l.svcCtx.Config.Email.From, l.svcCtx.Config.Email.Secret, l.svcCtx.Config.Email.Host), tlsConfig); err != nil {
	//     return nil, errors.Wrap(err, "发送邮件失败")
	// }

	return &types.Response{Code: 0, Msg: "邮件发送成功"}, nil
}
