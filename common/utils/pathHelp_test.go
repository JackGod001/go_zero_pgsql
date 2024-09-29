package utils

import (
	"testing"
)

// 获取项目根目录的绝对路径
func TestG(t *testing.T) {
	t.Log(GetRootPath())
	t.Log(GetConfigPath())

}
