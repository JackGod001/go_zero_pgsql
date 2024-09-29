package book_words

import (
	"context"
	"encoding/json"
	"fmt"
	"go_zero_pgsql/common/utils/redis_pika"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	SortSetKey = "review_word"
	userId     = "1"
	bookId     = "1"
)

// 在测试之前设置测试所需的数据
func setUp() *redis.Client {
	rdb := redis_pika.NewRedisClient()

	for i := 1; i < 3; i++ {
		word := ReviewWord{
			ReviewCount: "0",
			WordId:      strconv.Itoa(i),
			CreateAt:    time.Now(),
			NextReview:  time.Now(),
		}
		_, err := ReviewWordUpdate(rdb, userId, bookId, word)
		if err != nil {
			panic("初始化测试数据失败：" + err.Error())
		}

	}

	return rdb
}

// 清理测试数据
func tearDown(rdb *redis.Client) {
	// 清理测试数据
	//rdb := redis_pika.NewRedisClient()
	////删除集合中的元素
	rdb.Del(context.Background(), SortSetKey)
	rdb.Close()
}

// EbbinghausForgettingCurve 根据艾宾浩斯遗忘曲线计算下次复习的时间。
// 它根据复习的次数来计算间隔，我们定义复习时间节点如下：
// [5分钟, 30分钟, 12小时, 1天, 2天, 4天, 7天, 15天]
func Test_EbbinghausForgettingCurve(t *testing.T) {
	t.Log("测试艾宾浩斯遗忘曲线")
	reviewCount := 0
	t.Log(reviewCount, "当前次数")
	//根据复习次数获取时间间隔
	timeInterval := EbbinghausForgettingCurve(reviewCount)
	// 打印时间间隔
	t.Log(timeInterval, "下次时间间隔")
	// 当前时间
	timeNow := time.Now()
	//打印当前时间
	t.Log(timeNow, "当前时间")
	//当前时间加上时间间隔
	timeOut := timeNow.Add(timeInterval).Format(time.RFC3339)
	t.Log(timeOut, "下次时间")
}

func Test_ReviewWordUpdate(t *testing.T) {
	t.Log("测试复习单词更新")
	rdb := setUp()
	defer tearDown(rdb)
}

func Test_GetWordsToReview_NoWords(t *testing.T) {
	t.Log("测试复习单词更新")
	rdb := setUp()
	defer tearDown(rdb)

	// 获取当前时间
	currentTime := time.Now()
	t.Log(currentTime, "当前时间")
	// 获取当前时间下达到时间需要复习的单词
	words, err := GetWordsToReview(rdb, SortSetKey, currentTime, 0)
	if err != nil {
		fmt.Println("获取单词错误:", err)
		return
	}

	if len(words) == 0 {
		t.Log("当前时间下没有需要复习的单词")
	}

}

// 测试复习时间在当前时间之前的情况
func Test_GetWordsToReview_PastReviewTime(t *testing.T) {
	rdb := setUp()
	defer tearDown(rdb)

	// 获取当前时间下需要复习的单词
	currentTime := time.Now().Add(500 * time.Hour)
	words, err := GetWordsToReview(rdb, SortSetKey, currentTime, 1)
	if err != nil {
		t.Errorf("获取单词错误: %v", err)
	}

	// 检查是否有需要复习的单词
	if len(words) == 0 {
		t.Error("没有获取到需要复习的单词")
	}
	// 获取第一个单词
	firstWordStr := words[0]
	var firstWordStrObject ReviewWord
	_ = json.Unmarshal([]byte(firstWordStr), &firstWordStrObject)
	//判断当前时间
	if currentTime.After(firstWordStrObject.NextReview) {
		// 更新下次复习时间以及,学习次数
		_, err := ReviewWordUpdate(rdb, userId, bookId, firstWordStrObject)
		if err != nil {
			t.Errorf("复习单词更新错误: %v", err)
		}

	}

}
