package redis_pika

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type SortedSetMember struct {
	Member string
	Score  float64
}

// 获取有序集合中的所有成员,按照分数从小到大
func GetSortSetAllMembersZRangeByScore(rdb *redis.Client, key string) ([]string, error) {
	// 获取有序集合中的所有成员，并按照分数从小到大排序
	items := make([]string, 0)
	members, err := rdb.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		items = append(items, member)
	}
	return items, nil
}

// 获取有序集合中的所有成员和对应的分数，按照分数从小到大排序
func GetSortSetAllMembersWithScores(rdb *redis.Client, key string) ([]SortedSetMember, error) {
	// 获取有序集合中的所有成员和对应的分数，并按照分数从小到大排序
	items := make([]SortedSetMember, 0)
	membersWithScores, err := rdb.ZRangeWithScores(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	for _, member := range membersWithScores {
		item := SortedSetMember{
			Member: member.Member.(string),
			Score:  member.Score,
		}
		items = append(items, item)
	}
	return items, nil
}

// 获取有序集合中的所有成员,不包含分数
func GetSortSetAllMembers(rdb *redis.Client, key string) ([]string, error) {
	// 获取有序集合中的所有成员，并按照分数从小到大排序
	items := make([]string, 0)
	members, err := rdb.ZRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		items = append(items, member)
	}
	return items, nil
}

// 添加单词到有序集合
func AddToSortSet(rdb *redis.Client, key string, score float64, member string) error {
	// 添加到有序集合中
	err := rdb.ZAdd(context.Background(), key, redis.Z{Score: score, Member: member}).Err()
	return err
}
