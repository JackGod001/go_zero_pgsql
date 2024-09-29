package utils

import (
	"os"
	"path/filepath"
)

// 获取项目根目录的绝对路径
func GetRootPath() string {
	//这是用在测试文件的
	path, _ := filepath.Abs("./")
	//当前路径的父路径
	path = filepath.Dir(path)
	//返回根路径
	path = filepath.Dir(path)
	return path
}
func GetExecutableRootPath() string {
	//执行时编译好的api文件 在根目录下 data/server/api 这里用来获取根目录
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	root := filepath.Dir(filepath.Dir(filepath.Dir(ex)))
	return root
}
func GetConfigPath() string {
	str := GetExecutableRootPath()
	return str + "/config"
}
