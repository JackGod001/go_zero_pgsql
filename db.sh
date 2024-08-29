#!/bin/bash

# 加载 .env 文件中的环境变量
if [ -f .env ]; then
    export $(cat .env | sed 's/#.*//g' | xargs)
fi

# 设置 PostgreSQL 连接参数
export PGHOST=localhost  # 如果需要，可以更改为特定的 IP 地址
export PGPORT=$POSTGRESQL_PORT       # 确保这与 docker-compose.yml 中暴露的端口一致

# 等待 PostgreSQL 启动
echo "Waiting for PostgreSQL to start..."
until PGPASSWORD=$DB_PASSWORD psql -h $PGHOST -U $DB_USER -d $DB_NAME -c '\q'; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

echo "PostgreSQL is up - executing command"

# 创建用户表并插入随机数据
PGPASSWORD=$DB_PASSWORD psql -h $PGHOST -U $DB_USER -d $DB_NAME << EOF
-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    age INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入随机数据
INSERT INTO users (username, email, age) VALUES
('user1', 'user1@example.com', 25),
('user2', 'user2@example.com', 30),
('user3', 'user3@example.com', 35),
('user4', 'user4@example.com', 28),
('user5', 'user5@example.com', 40);

-- 显示插入的数据
SELECT * FROM users;
EOF
