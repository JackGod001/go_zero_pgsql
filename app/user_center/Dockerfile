# 第一阶段：构建阶段
FROM golang:1.22.6-alpine AS builder

# 标记这个阶段的名称为 gobuilder
LABEL stage=gobuilder

# 设置环境变量
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

# 将 Alpine 的软件源更改为阿里云的镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 更新 Alpine 的包管理器并安装 tzdata（时区数据）
RUN apk update --no-cache && apk add --no-cache tzdata

# 设置工作目录为 /build
WORKDIR /build

# 复制 go.mod 和 go.sum 文件到容器的当前工作目录（/build）
ADD go.mod .
ADD go.sum .

# 下载依赖
RUN go mod download

# 复制所有源代码到容器的当前工作目录（/build）
# 注意源代码目录是docker-compose.yml文件中context指定的目录
COPY . .

# 复制配置文件到 /app/etc 目录
# 注意：这里假设主机上的 cmd/api/etc 目录包含必要的配置文件
COPY cmd/api/etc /app/etc

# 编译应用
# 源文件：cmd/api/user.go
# 输出：/app/user（容器内路径）
RUN go build -ldflags="-s -w" -o /app/user cmd/api/user.go

# 第二阶段：运行阶段
FROM scratch

# 从 builder 阶段复制 SSL 证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# 从 builder 阶段复制时区信息
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

# 设置时区
ENV TZ Asia/Shanghai

# 设置工作目录为 /app
WORKDIR /app

# 从 builder 阶段复制编译好的应用到 /app 目录
COPY --from=builder /app/user /app/user

# 从 builder 阶段复制配置文件到 /app/etc 目录
COPY --from=builder /app/etc /app/etc

# 设置容器启动命令
# ./user：运行编译好的应用
# -f etc/user-api.yaml：指定配置文件的路径（容器内路径）
CMD ["./user", "-f", "etc/user-api.yaml"]
