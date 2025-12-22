package svc

import (
	"github.com/dszqbsm/online-judge/api/internal/config"
	"github.com/dszqbsm/online-judge/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	LogInterceptor rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		LogInterceptor: middleware.NewLogInterceptorMiddleware().Handle,
	}
}
