package svc

import (
	"github.com/dszqbsm/online-judge/api/internal/config"
	"github.com/dszqbsm/online-judge/api/internal/middleware"
	"github.com/dszqbsm/online-judge/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	LogInterceptor rest.Middleware

	UserModel           model.UserModel
	PromblemModel       model.ProblemModel
	TestCaseModel       model.TestCaseModel
	SubmissionModel     model.SubmissionModel
	SubmissionCaseModel model.SubmissionCaseModel

	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:         c,
		LogInterceptor: middleware.NewLogInterceptorMiddleware().Handle,

		UserModel:           model.NewUserModel(conn),
		PromblemModel:       model.NewProblemModel(conn),
		TestCaseModel:       model.NewTestCaseModel(conn),
		SubmissionModel:     model.NewSubmissionModel(conn),
		SubmissionCaseModel: model.NewSubmissionCaseModel(conn),

		Redis: redis.MustNewRedis(c.Redis),
	}
}
