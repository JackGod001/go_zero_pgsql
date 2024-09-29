package yamlConf

import (
	"flag"
	"fmt"
	"go_zero_pgsql/common/utils"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Salt    string
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	Cache cache.CacheConf
	Redis redis.RedisConf
}

func GetEnvConfig(key string) string {
	// 获取项目根目录
	rootPath := utils.GetRootPath()

	// 构建完整的 .env 文件路径
	envFilePath := filepath.Join(rootPath, ".env")

	// 加载 .env 文件中的配置参数
	if err := godotenv.Load(envFilePath); err != nil {
		//打印错误
		fmt.Println("加载.env 失败")
		//结束程序
		os.Exit(1)
	}
	//加载配置参数
	fmt.Println("加载.env 成功")
	envConfig := os.Getenv(key)
	fmt.Println(envConfig)
	return envConfig
}
func GetYamlConf() Config {
	var rootPath = utils.GetRootPath()
	var yamlFile = rootPath + "/app/usercenter/cmd/api/etc/learnToUseUser.yaml"
	var configFile = flag.String("f", yamlFile, "the config file")
	flag.Parse()
	var c Config
	//获取根目录下的.env文件
	conf.MustLoad(*configFile, &c)
	c.Redis.Host = "localhost:" + GetEnvConfig("Redis_Expose_Port")

	return c
}
