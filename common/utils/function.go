package utils

import (
	"context"
	"crypto/md5"
	"fmt"
	"go_zero_pgsql/common/globalkey"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"

	"github.com/zeromicro/go-zero/core/logx"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// 常规 jwt 中设置的用户id 是int64
//func GetUserId(ctx context.Context) int64 {
//	var uid int64
//	if jsonUid, ok := ctx.Value(globalkey.SysJwtUserId).(json.Number); ok {
//		if int64Uid, err := jsonUid.Int64(); err == nil {
//			uid = int64Uid
//		} else {
//			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
//		}
//	}
//
//	return uid
//}

// casdoor id 是string uuid 36位长度的字符串
//
//	func GetUserId(ctx context.Context) string {
//		var uid string
//		if jsonUid, ok := ctx.Value(globalkey.SysJwtUserId).(string); ok {
//			uid = jsonUid
//		}
//		return uid
//	}
func GetCasdoorUserId(ctx context.Context) string {
	var uid string
	if jsonUid, ok := ctx.Value(globalkey.SysJwtUserId).(string); ok {
		uid = jsonUid
	}
	return uid
}

func GetCasdoorUser(ctx context.Context, CasdoorClient *casdoorsdk.Client) (*casdoorsdk.User, error) {
	user, err := CasdoorClient.GetUserByUserId(GetCasdoorUserId(ctx))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ArrayUniqueValue[T any](arr []T) []T {
	size := len(arr)
	result := make([]T, 0, size)
	temp := map[any]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}

	return result
}

func ArrayContainValue(arr []int64, search int64) bool {
	for _, v := range arr {
		if v == search {
			return true
		}
	}

	return false
}

func Intersect(slice1 []int64, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	n := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}

	return n
}

func Difference(slice1 []int64, slice2 []int64) []int64 {
	m := make(map[int64]int)
	n := make([]int64, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, v := range slice1 {
		times, _ := m[v]
		if times == 0 {
			n = append(n, v)
		}
	}

	return n
}
func LogErrorWithContext(ctx context.Context, appId int64, otherErrMsg string, err error) {
	logMsg := fmt.Sprintf("应用id: %d %s %s", appId, otherErrMsg, err.Error())
	logx.WithContext(ctx).Error(logMsg)
}
