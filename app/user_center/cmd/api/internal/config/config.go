package config

import (
	"go_zero_pgsql/common/i18n"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Database struct {
		PGDataSource string
	}
	I18nConf        i18n.Conf
	ProjectRootPath string
	Cache           cache.CacheConf
	Redis           redis.RedisConf
	Salt            string
	//go-zero内置的jwt认证，暂未使用
	//Auth          rest.AuthConf
	//使用casdoor的jwt认证
	CasdoorConfig casdoorsdk.AuthConfig
	ProjectConf   ProjectConf
	//McmsRpc       zrpc.RpcClientConf
}

type ProjectConf struct {
	EmailCaptchaExpiredTime int `json:",default=600"`
	//是否允许初始化数据库
	AllowInit bool `json:",default=true"`

	//SmsTemplateId           string `json:",optional"`
	//SmsAppId                string `json:",optional"`
	//SmsSignName             string `json:",optional"`
	//SmsParamsType           string `json:",default=json,options=[json,array]"`
	//RegisterVerify          string `json:",default=captcha,options=[disable,captcha,email,sms,sms_or_email]"`
	//LoginVerify             string `json:",default=captcha,options=[captcha,email,sms,sms_or_email,all]"`
	//ResetVerify             string `json:",default=email,options=[email,sms,sms_or_email]"`
	//RefreshTokenPeriod      int    `json:",optional,default=24"` // refresh token valid period, unit: hour | 刷新 token 的有效期，单位：小时
	//AccessTokenPeriod       int    `json:",optional,default=1"`  // access token valid period, unit: hour | 短期 token 的有效期，单位：小时
}
