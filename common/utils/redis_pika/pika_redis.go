package redis_pika

import (
	"context"
	"go_zero_pgsql/common/utils/yamlConf"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	rdb *redis.Client
)

func NewRedisClient() *redis.Client {
	config := yamlConf.GetYamlConf()
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Redis.Host, // use default Addr
			Password: config.Redis.Pass, // no password set
			DB:       0,                 // use default DB
		})
	}
	return rdb
}
