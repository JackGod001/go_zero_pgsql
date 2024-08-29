

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
wrk -t12 -c40000 -d30s -s test.lua http://localhost:8887

```