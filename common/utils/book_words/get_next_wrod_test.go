package book_words

import (
	"context"
	"encoding/json"
	"fmt"
	"go_zero_pgsql/common/globalkey"
	"go_zero_pgsql/common/utils/redis_pika"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thoas/go-funk"
)

var (
	ctx = context.Background()
)

func ClearAllKey(rdb *redis.Client) {
	// 清理测试数据
	//rdb := redis_pika.NewRedisClient()
	_, err := rdb.FlushAll(context.Background()).Result()
	if err != nil {
		fmt.Println("redis flushall error:", err)
		os.Exit(1)
	}
	rdb.Close()
}

// 初始化 Redis 客户端
func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:46378", // 设置 Redis 实例的地址
		Password: "G62m50oigInCsf",  // 没有设置密码
		DB:       0,                 // 默认数据库
	})
}

// // 获取要学习或要复习的单词
func GetNextWordToStudyOrReview(client *redis.Client, userId string) (string, error) {
	//根据用户ID获取当前学习的 book_id
	bookId, err := GetUserCurrentBookID(client, userId)
	if err != nil {
		return "", err
	}
	//检查用户是否完成了此课本的学习
	isComplete, err := IsBookStudyComplete(client, userId, bookId)
	if err != nil {
		return "", err
	}
	if isComplete {
		return "", err
	}

	//先获取需要复习的下一个单词
	nextWordID, err := GetNextReviewWordID(client, userId, bookId)
	if err != nil {
		return "", err
	}
	//将这个单词下次复习时间更新,如果没有下次复习时间，则将这个单词标记为已经完成,检查用户对应课本所有掌握的单词数量跟用户对应课本的单词数量是否一致,如果一致将用户对应课本是否完成标记上
	if nextWordID == "" {
		nextWord, err := GetNextStudyWordID(client, userId, bookId)
		if err != nil {
			return "", err
		}
		nextWordID = nextWord.Member
		if nextWordID != "" {
			//将新学单词对应的复习时间记录,方便下次检索到
			_, err := ReviewWordUpdate(
				client,
				userId,
				bookId,
				ReviewWord{WordId: nextWordID, CreateAt: time.Now(), NextReview: time.Now(), ReviewCount: ""},
			)
			if err != nil {
				return "", err
			}
		}
	}
	return nextWordID, nil
}

// 获取用户当前学习的 book_id
func GetUserCurrentBookID(client *redis.Client, userId string) (string, error) {
	// 实际情况中可能需要从用户数据中获取当前学习的书籍ID
	bookId, err := client.Get(ctx, globalkey.UserBookPrefix+userId).Result()
	if err != nil {
		return "", err
	}
	return bookId, nil
}

// 设置用户课本已经掌握的单词数量
func AddUserBookWordCount(client *redis.Client, userId string, bookId string) error {
	key := globalkey.UserBookWordCountPrefix + userId + ":" + bookId
	valStr, err := client.Get(ctx, key).Result()
	// 获取课本的单词总数量
	if err != nil {
		return nil
	}
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return nil
	}
	valInt++
	return client.Set(ctx, key, valInt, 0).Err()
}

// 根据用户ID和课本ID判断课本下的单词用户是否都已完成学习
func IsBookStudyComplete(client *redis.Client, userId string, bookId string) (bool, error) {
	// 获取已经掌握的单词数量
	key := globalkey.UserBookWordCountPrefix + userId + ":" + bookId
	valStr, err := client.Get(ctx, key).Result()
	// 获取课本的单词总数量
	if err != nil && err != redis.Nil {
		return false, err
	}
	set1, err := redis_pika.GetSortSetAllMembersWithScores(client, globalkey.BookWordPrefix+bookId)
	if err != nil {
		return false, err
	}
	if valStr == "" {
		valStr = "0"
	}
	//将字符串转化为int
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return false, err
	}
	return len(set1) == valInt, nil
}

// 获取已到达复习时间的单词ID
func GetNextReviewWordID(client *redis.Client, userId string, bookId string) (string, error) {
	// 根据用户id,课本id获取需要复习的下一个单词id
	currentTime := time.Now().Add(5 * time.Minute)
	//currentTime := time.Now()
	SortSetKey := globalkey.UserStudyBookPrefix + userId + ":" + bookId
	words, err := GetWordsToReview(client, SortSetKey, currentTime, 1)

	if err != nil {
		return "", err
	}
	if len(words) == 0 {

		return "", nil // 没有要复习的单词
	}
	// 获取第一个单词
	firstWordStr := words[0]
	var firstWordStrObject ReviewWord
	_ = json.Unmarshal([]byte(firstWordStr), &firstWordStrObject)
	//判断当前时间
	if currentTime.After(firstWordStrObject.NextReview) {
		// 更新下次复习时间以及,学习次数
		_, err := ReviewWordUpdate(client, userId, bookId, firstWordStrObject)
		if err != nil {
			return "", err
		}

	}
	return firstWordStrObject.WordId, nil
}

// 获取下一个要学习的单词ID
func GetNextStudyWordID(client *redis.Client, userId string, bookId string) (redis_pika.SortedSetMember, error) {
	//获取当前课本下的所有单词ID 有序集合
	SortSetKey1 := globalkey.BookWordPrefix + bookId
	set1, _ := redis_pika.GetSortSetAllMembersWithScores(client, SortSetKey1)
	//获取当前用户对应的课本已经在学习的单词ID 有序集合
	SortSetKey2 := globalkey.UserBookSortSetPrefix + userId + ":" + bookId
	set2, _ := redis_pika.GetSortSetAllMembersWithScores(client, SortSetKey2)

	//求在课本下,不在用户对应课本中的差集中第一个单词
	fmt.Println("set1:", set1)
	fmt.Println("set2:", set2)
	r1Interface, _ := funk.Difference(set1, set2)
	r1, ok := r1Interface.([]redis_pika.SortedSetMember)
	if !ok {
		fmt.Println("Failed to convert r1 to []string")
		return redis_pika.SortedSetMember{}, nil
	}
	if len(r1) < 1 {
		return redis_pika.SortedSetMember{}, nil
	}
	var firstWord redis_pika.SortedSetMember
	firstWord = r1[0]
	fmt.Println("r1:", r1)

	//将这个单词以及排序字段添加到用户对应课本中的单词 有序集合中 用于后续快速筛选需要复习单词
	_ = redis_pika.AddToSortSet(client, SortSetKey2, firstWord.Score, firstWord.Member)
	return firstWord, nil
}

//
//// 标记某些单词已经完全掌握
//func MarkWordsAsMastered(client *redis.Client, userId string, wordIDs []string, bookId string) error {
//	key := "user_master_words:" + userId + ":" + bookId
//	_, err := client.SAdd(ctx, key, wordIDs).Result()
//	return err
//}
//
//// 标记课本为已学习完成
//func MarkTextbookAsStudyComplete(client *redis.Client, userId string, bookId string) error {
//	key := "complete_book:" + userId + ":" + bookId
//	return client.Set(ctx, key, "true", 0).Err()
//}

//	func AddWord(client *redis.Client, word Word) error {
//		// 构造 Hash 表的键名
//		key := "words:" + word.UserID + ":" + word.TextbookID + ":" + word.WordID
//		//words:user123:book457:*
//		// 将额外字段转换为字符串
//		reviewCountStr := strconv.Itoa(word.ReviewCount)
//		createAtStr := word.CreateAt.Format(time.RFC3339)
//		nextReviewStr := word.NextReview.Format(time.RFC3339)
//
//		// 发起 HSET 命令来添加字段和值
//		return client.HMSet(ctx, key, "ReviewCount", reviewCountStr, "CreateAt", createAtStr, "NextReview", nextReviewStr).Err()
//	}
//
// 添加课本单词ID到有序集合中
func AddBookWordIdToBookSortSet(client *redis.Client, bookId string, word Word) error {
	key := globalkey.BookWordPrefix + bookId + ":" + word.ID
	return redis_pika.AddToSortSet(client, key, word.Order, word.ID)
}

// 设置课本:单词id=单词详细信息到hashmap中
//func AddWordToBookSortSet(client *redis.Client, bookId string, wordID string) error {
// 构造 Hash 表的键名
//key := globalkey.BookWordPrefix + bookId + ":" + wordID
//words:user123:book457:*
// 将
//}

// 更新用户对应课本中单词的复习时间
func ReviewWordUpdate(client *redis.Client, userId string, bookId string, word ReviewWord) (bool, error) {
	key := globalkey.UserStudyBookPrefix + userId + ":" + bookId
	err := DeleteWordByID(client, key, word)
	if err != nil {
		return false, err
	}

	//复习次数
	CreatAtStr := word.CreateAt.Format(time.RFC3339)
	//艾宾浩斯时间间隔
	ReviewCountInt, _ := strconv.Atoi(word.ReviewCount)
	timeInterval := EbbinghausForgettingCurve(ReviewCountInt)
	//如果返回的是 0
	if timeInterval == 0 {
		err = AddUserBookWordCount(client, userId, bookId)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	nextReview := word.CreateAt.Add(timeInterval)
	//下次复习时间=复习次数请求艾宾浩斯曲线获取时间间隔+创建时间
	nextReviewStr := nextReview.Format(time.RFC3339)
	//这里是复习次数+1 在计算完毕下次复习时间后
	reviewCountStr := strconv.Itoa(ReviewCountInt + 1)
	wordObject := map[string]string{
		"WordId":      word.WordId,
		"ReviewCount": reviewCountStr,
		"CreateAt":    CreatAtStr,
		"NextReview":  nextReviewStr,
	}
	//记录复习次数,创建时间,下次复习时间
	member, err := json.Marshal(wordObject)
	if err != nil {
		return false, err
	}
	// 将成员添加到有序集合中
	//分数:下次复习时间
	score := float64(nextReview.Unix())
	//打印用户id 课本id 添加的 member 学习记录
	fmt.Printf("用户id:%s  课本id:%s  添加的 member:%s  学习记录:%v", userId, bookId, member, wordObject)

	return false, client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

// 测试添加单词
func Test_AddWord(t *testing.T) {
	// 示例使用
	client := NewClient()
	//_ := client.FlushDB(ctx).Err()
	fmt.Println("清空数据库成功")

	addTestWord(client)

	set1, _ := redis_pika.GetSortSetAllMembersWithScores(client, globalkey.BookWordPrefix+bookId)
	fmt.Println("set1:", set1)

}

func addTestWord(client *redis.Client) bool {
	for i := 1; i <= 100; i++ {
		var word Word
		word.ID = strconv.Itoa(i)
		word.Order = float64(i)
		key := globalkey.BookWordPrefix + bookId
		errr := redis_pika.AddToSortSet(client, key, word.Order, word.ID)
		if errr != nil {
			fmt.Println("添加单词失败:", errr)
			return true
		}
		fmt.Println("添加单词成功第", i, "条单词")
	}
	return true
}

// 设置用户当前学习的课本id
func AddUserBookIdToUserBookSortSet(client *redis.Client, userId string, bookId string) error {
	key := globalkey.UserBookPrefix + userId
	return client.Set(ctx, key, bookId, 0).Err()
}

// 测试设置用户当前学习的课本id
func Test_AddUserBookIdToUserBookSortSet(t *testing.T) {
	client := NewClient()
	//defer ClearAllKey(client)
	userId := "1"
	bookId := "1"
	err := AddUserBookIdToUserBookSortSet(client, userId, bookId)
	if err != nil {
		fmt.Println("设置用户当前学习的课本id失败:", err)
		return
	}
	bookId, _ = GetUserCurrentBookID(client, userId)
	fmt.Println("bookId:", bookId)
	fmt.Println(userId, bookId, "设置用户当前学习的课本成功")
}

// 测试用户获取要学习或者要复习的单词
func Test_GetWordsToReview(t *testing.T) {
	client := NewClient()
	//defer ClearAllKey(client)

	err := AddUserBookIdToUserBookSortSet(client, userId, bookId)
	if err != nil {
		fmt.Println("设置用户当前学习的课本id失败:", err)
		return
	}
	addTestWord(client)

	userId := "1"
	wordId, err := GetNextWordToStudyOrReview(client, userId)
	//测试单词ID
	if err != nil {
		fmt.Println("获取单词失败:", err)
		return
	}
	fmt.Println("单词ID:", wordId)

}
