package redis_pika

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func PrintAllHashes(client *redis.Client, ctx context.Context) error {
	var cursor uint64 = 0
	var keys []string

	// 使用 SCAN 命令遍历所有键
	for {
		var err error
		keys, cursor, err = client.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return err
		}

		// 遍历当前扫描得到的键
		for _, key := range keys {
			// 使用 HGETALL 命令获取哈希表的字段和对应的值
			fields, err := client.HGetAll(ctx, key).Result()
			if err != nil {
				return err
			}

			// 打印键名
			fmt.Println("Hash table:", key)

			// 打印字段和值
			for field, value := range fields {
				fmt.Printf("  %s: %s\n", field, value)
			}
		}

		// 如果游标为0，则表示遍历完成
		if cursor == 0 {
			break
		}
	}

	return nil
}

func Test_main(t *testing.T) {
	client := NewRedisClient()
	err := PrintAllHashes(client, Ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
