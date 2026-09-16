// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero/zerogen-tmp/internal/config"
	"zero/zerogen-tmp/internal/middleware"
)

type ServiceContext struct {
	Config        config.Config
	CasbinHandler rest.Middleware
	JwtHandler    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CasbinHandler: middleware.NewCasbinHandlerMiddleware().Handle,
		JwtHandler:    middleware.NewJwtHandlerMiddleware().Handle,
	}
}
