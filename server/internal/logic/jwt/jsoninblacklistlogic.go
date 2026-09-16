package jwt

import (
	"context"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type JsonInBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonInBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JsonInBlacklistLogic {
	return &JsonInBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// JsonInBlacklist 将当前请求携带的 JWT 加入黑名单，使其立即失效。
// 原 Gin 接口无入参，token 取自请求头，前端不传请求体。
func (l *JsonInBlacklistLogic) JsonInBlacklist(token string) (resp *types.Response, err error) {
	if token == "" {
		return nil, errors.New("JWT令牌不能为空")
	}

	jwtRecord := model.SysJwtBlacklist{
		Jwt: token,
	}
	if err = l.svcCtx.DB.WithContext(l.ctx).Create(&jwtRecord).Error; err != nil {
		return nil, errors.Wrap(err, "JWT加入黑名单失败")
	}

	logx.WithContext(l.ctx).Infow("JWT加入黑名单成功",
		logx.Field("module", "jwt"),
		logx.Field("action", "json_in_blacklist"),
	)

	return &types.Response{Code: 0, Msg: "jwt作废成功"}, nil
}
