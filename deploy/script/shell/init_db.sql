-- 如果用户数据库存在，则删除
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_database WHERE datname = '${USER_DB_NAME}') THEN
        EXECUTE 'DROP DATABASE ${USER_DB_NAME}';
END IF;
END $$;

-- 创建用户数据库
CREATE DATABASE ${USER_DB_NAME};

-- 切换到用户数据库
\c ${USER_DB_NAME}

-- 如果用户表存在，则删除
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
        EXECUTE 'DROP TABLE users';
END IF;
END $$;

-- 创建用户表
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password_hash VARCHAR(100) NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 插入模拟用户数据
INSERT INTO users (username, email, password_hash) VALUES
                                                       ('user1', 'user1@example.com', 'hashed_password1'),
                                                       ('user2', 'user2@example.com', 'hashed_password2'),
                                                       ('user3', 'user3@example.com', 'hashed_password3');

-- 切换回 postgres 数据库
\c postgres

-- 如果订单数据库存在，则删除
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_database WHERE datname = '${ORDER_DB_NAME}') THEN
        EXECUTE 'DROP DATABASE ${ORDER_DB_NAME}';
END IF;
END $$;

-- 创建订单数据库
CREATE DATABASE ${ORDER_DB_NAME};

-- 切换到订单数据库
\c ${ORDER_DB_NAME}

-- 如果订单表存在，则删除
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'orders') THEN
        EXECUTE 'DROP TABLE orders';
END IF;
END $$;

-- 创建订单表
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER NOT NULL,
                        order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        total_amount DECIMAL(10, 2) NOT NULL,
                        status VARCHAR(20) NOT NULL,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 插入模拟订单数据
INSERT INTO orders (user_id, order_date, total_amount, status) VALUES
                                                                   (1, CURRENT_TIMESTAMP, 99.99, 'completed'),
                                                                   (2, CURRENT_TIMESTAMP, 49.99, 'pending'),
                                                                   (3, CURRENT_TIMESTAMP, 19.99, 'cancelled');
