# 生成api业务代码 ， 进入"服务/cmd/api/desc"目录下，执行下面命令
# 先cd 到"服务/cmd/api/desc"目录下
cd ../../../../
# 记录当前目录为一个变量
BASE_DIR=$(pwd)
appdir='user_center'
#输出变量
echo "项目根目录" $BASE_DIR
#设置api/desc目录为一个变量
API_DESC_DIR=$BASE_DIR/app/$appdir/cmd/api/desc
echo "api 目录: " $API_DESC_DIR

# 设置goctl模板路径
GOCTL_TEMPLATE_DIR=$BASE_DIR/deploy/goctl/1.7.1
echo "goctl 模板目录: "$GOCTL_TEMPLATE_DIR

#先前往api目录
cd  $API_DESC_DIR
# 执行命令 *.api -dir ../  --style=goZero -home=../../../../goctl/1.6.1
# goctl 这是在 go-zero 官方git master版本中的,1.6.2, (1.6.1 生成api时的引入公共文件时报错找不到包)
#$GOCTL_TEMPLATE_DIR/goctl api go -api  *.api -dir ../ -style=go --style=goZero --home=$GOCTL_TEMPLATE_DIR
$GOCTL_TEMPLATE_DIR/goctl api go -api  *.api -dir ../  -style=go --style=go_zero --home=$GOCTL_TEMPLATE_DIR


# 生成 swagger json 文件
SWAGGERAPI_DIR=$BASE_DIR/deploy/swaggerapi
#如果目录不存在则创建
mkdir -p $SWAGGERAPI_DIR
for app_dir in $BASE_DIR/app/*/; do
    app_name=$(basename "$app_dir")
    echo "Processing $app_name"
    api_file=$app_dir"cmd/api/desc/main.api"
    echo "Checking API file for $api_file"
    if [ -f "$api_file" ]; then
        echo "Generating Swagger JSON for $app_name"
        $GOCTL_TEMPLATE_DIR/goctl api plugin -plugin goctl-swagger="swagger -filename $app_name.json" -api "$api_file" -dir $SWAGGERAPI_DIR
    fi
done
