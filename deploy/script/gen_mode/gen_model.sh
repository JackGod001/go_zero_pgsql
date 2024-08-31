#!/usr/bin/env bash

# 使用方法：
# ./gen_model.sh
# 生成的文件将在 ./gen_model 目录下

# 设置默认值
tables="*"
modeldir="./gen_model"

# 获取项目根目录路径
root_dir=$(git rev-parse --show-toplevel)
if [ $? -ne 0 ]; then
    echo "错误：无法确定项目根目录。请确保您在 Git 仓库中执行此脚本。"
    exit 1
fi

# 读取 .env 文件
if [ -f "$root_dir/.env" ]; then
    source "$root_dir/.env"
else
    echo "错误：在项目根目录中找不到 .env 文件。"
    exit 1
fi

# 检查必要的环境变量
required_vars=("APP_NAME" "DB_USER" "DB_PASSWORD"  "POSTGRESQL_EXPOSE_PORT")
for var in "${required_vars[@]}"; do
    if [ -z "${!var}" ]; then
        echo "错误：.env 文件中缺少必要的配置：$var"
        exit 1
    fi
done

# 要生成哪个数据库的代码
dbname=$USER_DB_NAME

# 数据库用户名
username=$DB_USER
# 数据库密码
passwd=$DB_PASSWORD
# 这里是docker-compose.yml中设置的服务名
host=${POSTGRES_SERVICE}
#.env文件中配置的容器对外端口
port=${POSTGRESQL_EXPOSE_PORT}


pwd
# 确保 gen_model 目录存在
mkdir -p "$modeldir"

# 获取 goctl 模板路径
currentPath=$(pwd)
goctlPath=$(realpath ../../goctl/1.7.1)
if [ ! -d "$goctlPath" ]; then
    echo "错误：找不到 goctl 目录：$goctlPath"
    exit 1
fi

# 设置 goctl 模板路径
export GOCTL_HOME="$goctlPath"
echo "goctl 模板目录: $GOCTL_HOME"

cd "$currentPath" || exit
echo "开始为数据库 $dbname 生成所有表的模型"

# 协议://用户名:密码@主机名:端口/数据库名?sslmode=disable
CON="postgres://${username}:${passwd}@localhost:${port}/${dbname}?sslmode=disable"
echo "数据库连接信息：$CON"
# 执行 goctl 命令
$GOCTL_HOME/ model pg datasource -url=${CON} -table="${tables}" -dir="${modeldir}" -cache=true --style=goZero

if [ $? -ne 0 ]; then
    echo "错误：goctl 命令执行失败"
    exit 1
fi

echo "模型生成完成。生成的文件位于 $modeldir 目录。"
echo "请将生成的文件移动到相应服务的 model 目录，并记得修改 package 声明。"
