package book_words

import (
	"go_zero_pgsql/common/utils/redis_pika"

	"github.com/redis/go-redis/v9"
)

type Word struct {
	ID    string
	Order float64
}

// 添加课本对应的单词 到有序集合中
func BookAdd(rdb *redis.Client) string {
	// 设置课本名称
	key := "课本"
	// 添加单词到有序集合
	word1 := Word{ID: "单词ID1", Order: 1}
	_ = redis_pika.AddToSortSet(rdb, key, word1.Order, word1.ID)
	word2 := Word{ID: "单词ID2", Order: 2}
	_ = redis_pika.AddToSortSet(rdb, key, word2.Order, word2.ID)
	word30 := Word{ID: "单词ID30", Order: 30}
	_ = redis_pika.AddToSortSet(rdb, key, word30.Order, word30.ID)
	word31 := Word{ID: "单词ID31", Order: 30}
	_ = redis_pika.AddToSortSet(rdb, key, word31.Order, word31.ID)
	return key
}

// 添加用户对应的课本单词 到有序集合中
func UserBookAdd(rdb *redis.Client) string {
	// 设置用户课本名称
	key := "user_id_1_book_1"
	// 添加单词到有序集合
	word1 := Word{ID: "单词ID1", Order: 1}
	_ = redis_pika.AddToSortSet(rdb, key, word1.Order, word1.ID)
	word30 := Word{ID: "单词ID30", Order: 30}
	_ = redis_pika.AddToSortSet(rdb, key, word30.Order, word30.ID)
	return key
}
