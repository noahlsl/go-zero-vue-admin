// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"zero/internal/config"
	"zero/internal/handler"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 业务状态码，与原 Gin 后端 response.Response 保持一致
const (
	// successCode 业务成功
	successCode = 0
	// errorCode 业务失败
	errorCode = 7
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	initResponseHandler()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// initResponseHandler 统一响应格式为 {code,data,msg}，与原 Gin 后端保持一致，
// 使前端无需改动即可切换后端实现。
func initResponseHandler() {
	httpx.SetOkHandler(func(_ context.Context, v any) any {
		if resp, ok := v.(*types.Response); ok {
			if resp == nil {
				return &types.Response{Code: successCode, Data: map[string]any{}, Msg: "成功"}
			}
			// 与原 Gin 后端 OkWithMessage/Ok() 保持一致：data 为空时返回空对象而非 null
			if resp.Data == nil {
				resp.Data = map[string]any{}
			}
			if resp.Msg == "" {
				resp.Msg = "成功"
			}
			return resp
		}
		return &types.Response{Code: successCode, Data: v, Msg: "成功"}
	})

	httpx.SetErrorHandlerCtx(func(_ context.Context, err error) (int, any) {
		return http.StatusOK, &types.Response{
			Code: errorCode,
			Data: map[string]any{},
			Msg:  err.Error(),
		}
	})
}
