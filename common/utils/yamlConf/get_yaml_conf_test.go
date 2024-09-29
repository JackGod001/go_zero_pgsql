package yamlConf

import (
	"testing"
)

func TestGetYamlConf(t *testing.T) {
	c := GetYamlConf()
	// 打印
	t.Log(c)
}
