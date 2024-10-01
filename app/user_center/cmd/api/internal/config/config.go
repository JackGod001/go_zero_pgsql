package config

import (
	"go_zero_pgsql/common/i18n"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Database struct {
		PGDataSource string
	}
	I18nConf i18n.Conf
	Cache    cache.CacheConf
	Redis    redis.RedisConf
	Salt     string
	Auth     struct {
		AccessSecret string
		AccessExpire int64
	}
}
