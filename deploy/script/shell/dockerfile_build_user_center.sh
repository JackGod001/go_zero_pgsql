#!/bin/bash

# 获取 Git 仓库的根目录,必须要有.git目录
ROOT_DIR=$(git rev-parse --show-toplevel)

echo "项目根目录: $ROOT_DIR"

# 使用 ROOT_DIR 变量进行后续操作
GOCTLDIR=$ROOT_DIR/deploy/goctl/1.7.1
echo "生成dockerfile文件的goctl目录: $GOCTLDIR"

# 某个服务的go.mod目录,注意是服务文件夹内的go.mod文件 这里不同服务修改不同的目录
GO_MOD_DIR=$ROOT_DIR/app/user_center
echo "进入服务的go.mod目录: $GO_MOD_DIR"
cd $GO_MOD_DIR
pwd
echo "开始生成dockerfile文件"

# 这里注意修改不同服务的go文件
$GOCTLDIR/goctl docker -go ./cmd/api/user.go
