#!/bin/bash

# 获取项目根目录路径
root_dir=$(git rev-parse --show-toplevel)
if [ $? -ne 0 ]; then
    echo "错误：无法确定项目根目录。请确保您在 Git 仓库中执行此脚本。"
    exit 1
fi

echo "项目根目录: $root_dir"

# 加载 .env 文件中的环境变量，如果文件不存在则报错
envFile=$root_dir/.env
if [ ! -f "$envFile" ]; then
    echo "错误：.env 文件未找到。"
    exit 1
fi

export $(cat $envFile | sed 's/#.*//g' | xargs)

# 设置 PostgreSQL 连接参数
export PGHOST=localhost  # 如果需要，可以更改为特定的 IP 地址
export PGPORT=$POSTGRESQL_EXPOSE_PORT       # 确保这与 docker-compose.yml 中暴露的端口一致

# 等待 PostgreSQL 启动
echo "等待 PostgreSQL 启动..."
until PGPASSWORD=$DB_PASSWORD psql -h $PGHOST -U $DB_USER -d $POSTGRES_SERVICE -c '\q'; do
  echo "PostgreSQL 尚未就绪 - 等待中"
  sleep 1
done

echo "PostgreSQL 已就绪 - 执行命令"

# 使用 envsubst 替换 SQL 文件中的环境变量，然后执行
echo "初始化数据库..."
envsubst < ./init_db.sql | PGPASSWORD=$DB_PASSWORD psql -h $PGHOST -U $DB_USER -d $POSTGRES_SERVICE

# 验证数据库和表的创建
echo "验证数据库和表的创建..."

# 检查 user_db
echo "检查用户数据库..."
PGPASSWORD=$DB_PASSWORD psql -h $PGHOST -U $DB_USER -d $USER_DB_NAME << EOF
\dt
SELECT COUNT(*) FROM users;
EOF

# 检查 order_db
echo "检查订单数据库..."
PGPASSWORD=$DB_PASSWORD psql -h $PGHOST -U $DB_USER -d $ORDER_DB_NAME << EOF
\dt
SELECT COUNT(*) FROM orders;
EOF

echo "脚本执行完成。"
