#!/bin/sh
echo "start entrypoint.sh"

# 接收传入的参数
PORT=$1
EXECUTABLE_PATH=$2
CONFIG_PATH=$3


echo "Executable path: $EXECUTABLE_PATH"
echo "Config path: $CONFIG_PATH"

# 确保可执行文件存在
if [ ! -f "$EXECUTABLE_PATH" ]; then
    echo "Error: Executable file not found at $EXECUTABLE_PATH"
    exit 1
fi


# 假设 $CONFIG_PATH 和 $EXECUTABLE_PATH 已经被设置

DLV_CMD="dlv --headless --listen=:$PORT --accept-multiclient --api-version=2 exec $EXECUTABLE_PATH"

# 检查 $CONFIG_PATH 是否为空
if [ -z "$CONFIG_PATH" ]; then
    echo "Notice: CONFIG_PATH is empty, running dlv without config file."
else
    # 如果 $CONFIG_PATH 不为空，检查配置文件是否存在
    if [ ! -f "$CONFIG_PATH" ]; then
        echo "Error: Config file not found at $CONFIG_PATH"
        exit 1
    fi
    # 配置文件存在，拼接完整的 dlv 命令
    DLV_CMD="$DLV_CMD -- -f $CONFIG_PATH"
fi

# 运行 dlv 命令
echo "Running dlv command... $DLV_CMD"
#$DLV_CMD >> debug.log 2>&1
$DLV_CMD
