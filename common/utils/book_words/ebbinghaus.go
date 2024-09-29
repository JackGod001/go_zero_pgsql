package book_words

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

//

// EbbinghausForgettingCurve 根据艾宾浩斯遗忘曲线计算下次复习的时间。
// 它根据复习的次数来计算间隔，我们定义复习时间节点如下：
// [5分钟, 30分钟, 12小时, 1天, 2天, 4天, 7天, 15天]
func EbbinghausForgettingCurve(reviewCount int) time.Duration {
	intervals := []time.Duration{
		5 * time.Minute,
		30 * time.Minute,
		12 * time.Hour,
		24 * time.Hour,
		2 * 24 * time.Hour,
		4 * 24 * time.Hour,
		7 * 24 * time.Hour,
		15 * 24 * time.Hour,
	}
	//intervals
	// 如果复习次数超出我们定义的间隔数，则认为这个单词已经被掌握
	if reviewCount >= len(intervals) {
		return 0
	}
	return intervals[reviewCount]
}

type ReviewWord struct {
	CreateAt    time.Time
	NextReview  time.Time
	ReviewCount string
	WordId      string
}

// DeleteWordByID 从有序集合中删除指定 WordId 的单词成员
func DeleteWordByID(client *redis.Client, key string, word ReviewWord) error {
	// 执行 ZREM 命令删除指定 WordId 的成员
	wordstr, _ := json.Marshal(word)
	_, err := client.ZRem(context.Background(), key, wordstr).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetWordsToReview(client *redis.Client, key string, currentTime time.Time, count int64) ([]string, error) {
	// 获取当前时间到未来某个时间点的单词
	// 未来时间点为当前时间
	words, err := client.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{
		Min:   "0",                                       // 分数的最小值，表示当前时间
		Max:   strconv.FormatInt(currentTime.Unix(), 10), // 分数的最大值，表示当前时间
		Count: count,
	}).Result()
	if err != nil {
		return nil, err
	}
	return words, nil
}
