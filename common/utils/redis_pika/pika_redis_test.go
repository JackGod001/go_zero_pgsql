package redis_pika

import (
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestConnet(t *testing.T) {
	rdb := NewRedisClient()

	err := rdb.Set(Ctx, "key", "value111", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(Ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(Ctx, "missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
	//

}
