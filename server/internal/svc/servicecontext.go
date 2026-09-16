package svc

import (
	"zero/internal/config"
	"zero/internal/dao/repos"
	"zero/internal/middleware"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	JwtHandler      rest.Middleware
	CasbinHandler   rest.Middleware
	OperationLogger rest.Middleware
	Cors            rest.Middleware
	DB              *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := repos.InitMySQL(
		c.MySQL.DSN,
		c.MySQL.MaxIdleConns,
		c.MySQL.MaxOpenConns,
		c.MySQL.ConnMaxLifetime,
	)
	if err != nil {
		logx.Must(err)
	}

	return &ServiceContext{
		Config:          c,
		JwtHandler:      middleware.NewJwtMiddleware(c.Auth.AccessSecret, c.Auth.AccessExpire, db),
		CasbinHandler:   middleware.NewCasbinHandlerMiddleware().Handle,
		OperationLogger: middleware.NewOperationLoggerMiddleware(db).Handle,
		Cors:            middleware.NewCorsMiddleware(c.Cors).Handle,
		DB:              db,
	}
}
