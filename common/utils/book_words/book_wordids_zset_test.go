package book_words

import (
	"fmt"
	"go_zero_pgsql/common/utils/redis_pika"
	"testing"

	"github.com/thoas/go-funk"
	"k8s.io/apimachinery/pkg/util/sets"
)

func Test_Book_Add(t *testing.T) {
	// 连接 Redis
	rdb := redis_pika.NewRedisClient()

	book := BookAdd(rdb)
	// 获取按顺序的所有单词ID
	words, err := redis_pika.GetSortSetAllMembers(rdb, book)
	if err != nil {
		fmt.Println("Failed to get words ordered:", err)
		return
	}

	// 打印结果
	fmt.Printf("按顺序的所有单词ID：\n")
	for _, word := range words {
		fmt.Println(word)
	}
}

func Test_UserBook_Add(t *testing.T) {
	// 连接 Redis
	rdb := redis_pika.NewRedisClient()
	book := UserBookAdd(rdb)
	// 获取按顺序的所有单词ID
	words, err := redis_pika.GetSortSetAllMembers(rdb, book)
	if err != nil {
		fmt.Println("Failed to get words ordered:", err)
		return
	}

	// 打印结果
	fmt.Printf("按顺序的所有单词ID：\n")
	for _, word := range words {
		fmt.Println(word)
	}
}

// 获取两个有序集合的差集 不包含排序字段
func Test_DifferenceBetweenSets(t *testing.T) {
	// Connect to Redis
	rdb := redis_pika.NewRedisClient()
	//清空所有keys
	err := rdb.FlushDB(redis_pika.Ctx).Err()
	if err != nil {
		fmt.Println("Failed to flush db:", err)
		return
	}
	set1key := BookAdd(rdb)
	set2key := UserBookAdd(rdb)
	set1, _ := redis_pika.GetSortSetAllMembersZRangeByScore(rdb, set1key)
	set2, _ := redis_pika.GetSortSetAllMembersZRangeByScore(rdb, set2key)
	fmt.Println("set1:", set1)
	fmt.Println("set2:", set2)
	// 使用 sets 包中的方法计算在 set1 中但不在 set2 中的单词
	difference := sets.NewString()
	difference.Insert(set1...)
	difference.Delete(set2...)
	fmt.Println("Difference between set1 and set2:", difference.List())
}

// 获取两个有序集合的差集
func Test_diff_funk(t *testing.T) {
	// Connect to Redis
	rdb := redis_pika.NewRedisClient()
	//清空所有keys
	err := rdb.FlushDB(redis_pika.Ctx).Err()
	if err != nil {
		fmt.Println("Failed to flush db:", err)
		return
	}
	set1key := BookAdd(rdb)
	set2key := UserBookAdd(rdb)
	set1, _ := redis_pika.GetSortSetAllMembersWithScores(rdb, set1key)
	set2, _ := redis_pika.GetSortSetAllMembersWithScores(rdb, set2key)
	fmt.Println("set1:", set1)
	fmt.Println("set2:", set2)
	r1Interface, _ := funk.Difference(set1, set2)
	r1, ok := r1Interface.([]redis_pika.SortedSetMember)
	if !ok {
		fmt.Println("Failed to convert r1 to []string")
		return
	}
	fmt.Println("r1 在 set1 中没在 set2 中", r1)
	if len(r1) > 0 {
		fmt.Println("r1 在 set1 中没在 set2 中", r1[0])
	}

}
