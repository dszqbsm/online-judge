package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MySQL   sqlx.SqlConf
	Redis   redis.RedisConf
	JwtAuth JwtAuthConf
}

type JwtAuthConf struct {
	AccessSecret string
	AccessExpire int64
}
