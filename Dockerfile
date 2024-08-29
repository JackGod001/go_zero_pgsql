FROM golang:1.22.6-alpine

# 设置环境变量
ENV GOPROXY https://goproxy.cn,direct
ENV TZ Asia/Shanghai

# 安装所需的包和设置时区
RUN apk update --no-cache && \
    apk add --no-cache tzdata git && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ >/etc/timezone

# 安装 modd 和 Delve
RUN go install github.com/cortesi/modd/cmd/modd@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest


# 设置工作目录
#WORKDIR /go
