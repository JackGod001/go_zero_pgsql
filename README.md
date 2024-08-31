
## 技术架构
1. postgres 集群监控 Pigsty
2. modd(热加载) 和 Delve(调试)
3. redis todo pika
## 目录架构
todo 关于model可以放到rpc的目录internal目录中,这样api就只能通过rpc->model
优点: 防止api和model耦合
缺点: 灵活性有所降低

## postgres 配置
context 文件需要导入驱动
```go
	_ "github.com/jackc/pgx/v5/stdlib"
```
然后执行 
```shell
go mod tidy
```
[service_context.go](app%2Fuser_center%2Fcmd%2Fapi%2Finternal%2Fsvc%2Fservice_context.go)
[user-api.yaml](app%2Fuser_center%2Fcmd%2Fapi%2Fetc%2Fuser-api.yaml)



并发测试
```shell
wrk -t12 -c4000 -d3s -s test.lua http://localhost:8887

```


## docker
### 重新构建
发布环境注意重新构建 [ROCKER-REBUILD.md](ROCKER-REBUILD.md)
### 不同服务的dockerfile以及go.mod
不同服务的目录下放不同的 go.mod ,这样有些服务不需要的依赖包就不会被下载
### 生成dockerfile
使用脚本
[dockerfile_build_user_center.sh](deploy%2Fscript%2Fshell%2Fdockerfile_build_user_center.sh)  
生成dockerfile使用的是1.7.1未发布版本 [goctl](deploy%2Fgoctl%2F1.7.1%2Fgoctl)  
在go.mod目录中执行!!! (shell已经有到go.mod)
### 1.7.1 goctl 自行修改的内容
自行修正了并生成的可执行文件  
解决不同目录下执行生成 dockerfile 文件中没有找到对应etc目录,如果使用官方的无法找到etc目录无法找到yml配置文件
copy对应的目录也不对
#如果使用官方的要自己找到go.mod 对应的etc目录以及配置文件,并添加如下代码结合文件查看代码添加在哪里

[Dockerfile](app%2Fuser_center%2FDockerfile)
```shell
# 检查后面的目录,是 doccker-compose context的目录 下的目录
RUN go build -ldflags="-s -w" -o /app/user cmd/api/user.go

# 左侧的路径是 go.mod 文件的相对路径 右侧是容器中的etc目录
COPY cmd/api/etc /app/etc
COPY --from=builder /app/etc /app/etc
# 启动命令 以容器中的配置文件启动的,目录是容器内的目录
CMD ["./user", "-f", "etc/user-api.yaml"]
```

## go-zero 生成代码的模板文件暂时采用 goctl 1.7.1
goctl可执行文件是根据官方master分支修改的,已经是1.7.1版本了,虽然go.mod是1.7.0的版本,但是生成的dockerfile是1.7.1的版本,所以可以不用修改
但是这里先采用1.7.1版本,此框架代码写时1.7.1版本还未发布
牵扯文件
[gen_api.sh](deploy%2Fscript%2Fgencode%2Fgen_api%2Fgen_api.sh)

## todo
postgres 集群监控 Pigsty
[哔哩哔哩](https://www.bilibili.com/video/BV13q4y1o74M/?spm_id_from=333.880.my_history.page.click&vd_source=ca29f7158bd0ff443c7d38352c028de4)





## 环境
### 开发
开发时注意不要在对应服务目录下生成go.mod,否则docker compose up 会报错,
在发布环境时,需要重新生成go.mod
todo: 对应服务下的 dockerfile 修改配置不同环境,docker-compose-dev.yml这样指向同一个服务的dockerfile
在gotctl初次生成的dockerfile上修改的 
```shell
# 启动执行  指定 docker-compose-dev.yml
docker compose -f docker-compose-dev.yml up
# 停止 
docker compose -f docker-compose-dev.yml down
# 重新build 不使用缓存
docker compose -f docker-compose-dev.yml build --no-cache
```

### 线上
```shell
#删除
docker compose down
## 重新build 不使用缓存
docker compose build --no-cache
# 启动
docker compose up -d   
```
