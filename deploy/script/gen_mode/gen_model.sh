#!/usr/bin/env bash

# 使用方法：
# ./gen_mode.sh
# 生成的文件将在 ./gen_mode 目录下

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
if [ -z "$APP_NAME" ] || [ -z "$DB_USER" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_NAME" ]; then
    echo "错误：.env 文件中缺少必要的数据库配置。"
    exit 1
fi

# 设置数据库连接参数
dbname=$DB_NAME
username=$DB_USER
passwd=$DB_PASSWORD
port=5432

# 确保 gen_mode 目录存在
mkdir -p "$modeldir"


currentPath=$(pwd)
cd ../../goctl/1.7.0
# 设置goctl模板路径
GOCTL_TEMPLATE_DIR=$(pwd)
echo "goctl 模板目录: "$GOCTL_TEMPLATE_DIR

cd  $currentPath
echo "开始为数据库 $dbname 生成所有表的模型"

#echo "开始创建库：$dbname 的表：$2"
$GOCTL_TEMPLATE_DIR/goctl model pg datasource -url="postgres://${username}:${passwd}@localhost:${port}/${dbname}?sslmode=disable" -table="${tables}" -dir="${modeldir}" -cache=true --style=goZero
echo "模型生成完成。生成的文件位于 $modeldir 目录。"
echo "请将生成的文件移动到相应服务的 model 目录，并记得修改 package 声明。"

