package base

import (
	"context"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// store 验证码存储，内存模式；多实例部署时替换为 Redis 存储
var store = base64Captcha.DefaultMemStore

// 验证码默认参数
const (
	captchaImgHeight = 80  // 验证码图片高度
	captchaImgWidth  = 240 // 验证码图片宽度
	captchaKeyLong   = 6   // 验证码位数
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Captcha 生成验证码，返回 captchaId 和 base64 图片。
// 原 Gin 接口无入参，前端不传请求体。
func (l *CaptchaLogic) Captcha() (resp *types.Captcha, err error) {
	// 使用数字验证码 driver
	driver := base64Captcha.NewDriverDigit(captchaImgHeight, captchaImgWidth, captchaKeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)

	id, b64s, _, err := cp.Generate()
	if err != nil {
		logx.Errorw("验证码生成失败", logx.Field("error", err.Error()))
		return nil, errors.Wrap(err, "验证码生成失败")
	}

	return &types.Captcha{
		CaptchaId: id,
		PicPath:   b64s,
		// 登录逻辑始终校验验证码，等价于原 Gin 配置 open-captcha: 0（一直开启），
		// 前端据此展示验证码输入框，因此恒为 true
		CaptchaLength: captchaKeyLong,
		OpenCaptcha:   true,
	}, nil
}
