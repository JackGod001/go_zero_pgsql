package errorx

var errorMsg map[int]string

const (
	ServerErrorCode         = 1000
	ParamErrorCode          = 1001
	CaptchaErrorCode        = 1002
	AccountErrorCode        = 1003
	PasswordErrorCode       = 1004
	NotPermMenuErrorCode    = 1005
	AddUserErrorCode        = 1016
	AccountDisableErrorCode = 1021
	AddConfigErrorCode      = 1024
	AddDictionaryErrorCode  = 1025
	AuthErrorCode           = 401
	ForbiddenErrorCode      = 1030
	UserIdErrorCode         = 1044

	// 1300 用户不可用
	UserDisableErrorCode = 1300

	OrderRepeatErrorCode = 2000
)

func init() {
	errorMsg = make(map[int]string)
	errorMsg[ServerErrorCode] = "服务繁忙，请稍后重试"
	errorMsg[CaptchaErrorCode] = "验证码错误"
	errorMsg[AccountErrorCode] = "账号错误"
	errorMsg[PasswordErrorCode] = "密码错误"
	errorMsg[NotPermMenuErrorCode] = "权限不足"
	errorMsg[AddUserErrorCode] = "账号已存在"
	errorMsg[AccountDisableErrorCode] = "账号已禁用"
	errorMsg[AddConfigErrorCode] = "配置已存在"
	errorMsg[AddDictionaryErrorCode] = "字典已存在"
	errorMsg[AuthErrorCode] = "授权已失效，请重新登录"
	errorMsg[ForbiddenErrorCode] = "禁止操作"
	errorMsg[UserIdErrorCode] = "用户不存在"
	errorMsg[UserDisableErrorCode] = "用户已禁用"
	errorMsg[OrderRepeatErrorCode] = "订单重复"

}

func MapErrMsg(errCode int) string {
	if msg, ok := errorMsg[errCode]; ok {
		return msg
	} else {
		return "服务繁忙，请稍后重试"
	}
}
