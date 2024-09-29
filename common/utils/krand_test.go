package utils

import (
	"fmt"
	"testing"

	"github.com/JackGod001/goRand"
)

func TestMd5ByString(t *testing.T) {
	//query = SqlBuildNormalData(query)
	slat := "Kdi8mTfc5sTXO7OG"
	s := Md5ByString("123123" + slat)
	t.Log(s)
}

// 测试Krand.go文件中的
// 测试随机字符串 KrandS 方法
func Test(t *testing.T) {
	fmt.Println(goRand.GRands(10, goRand.KcRandKindNum))
}
