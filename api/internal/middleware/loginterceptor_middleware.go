package middleware

import "github.com/dszqbsm/logx"

func NewLogInterceptorMiddleware() *logx.InterceptorMiddleware {
	return logx.NewInterceptorMiddleware()
}
