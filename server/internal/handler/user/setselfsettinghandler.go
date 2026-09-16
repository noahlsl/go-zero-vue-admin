package user

import (
	"io"
	"net/http"

	"zero/internal/logic/user"
	"zero/internal/svc"

	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetSelfSettingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer r.Body.Close()

		var setting map[string]any
		if err = sonic.Unmarshal(body, &setting); err != nil {
			logx.WithContext(r.Context()).Errorw("解析用户配置失败",
				logx.Field("error", err),
				logx.Field("module", "user"),
				logx.Field("action", "set_self_setting"),
			)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSetSelfSettingLogic(r.Context(), svcCtx)
		if err = l.SetSelfSetting(setting); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
