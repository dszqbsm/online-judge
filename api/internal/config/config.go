package config

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MySQL sqlx.SqlConf
}
