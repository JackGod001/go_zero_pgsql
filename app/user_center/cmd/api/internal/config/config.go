package config

import (
	"go_zero_pgsql/common/i18n"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Database struct {
		PGDataSource string
	}
	I18nConf    i18n.Conf
	ProjectRoot string
	Cache       cache.CacheConf
	Redis       redis.RedisConf
	Salt        string
	//go-zero内置的jwt认证，暂未使用
	//Auth          rest.AuthConf
	//使用casdoor的jwt认证
	CasdoorConfig casdoorsdk.AuthConfig
}
