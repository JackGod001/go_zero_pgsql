package request

import (
	"net/http"
)

func GetRequestDomain(r *http.Request) string {
	url := ""
	if r.TLS == nil {
		url = "http://"
	} else {
		url = "https://"
	}
	return url + r.Host
}

// GetRequestLang 从请求中获取语言设置
func GetRequestLang(r *http.Request) string {
	// 首先检查URL参数
	lang := r.FormValue("lang")
	if lang != "" {
		return lang
	}

	// 如果URL参数中没有，则检查Accept-Language头
	lang = r.Header.Get("Accept-Language")
	if lang != "" {
		return lang
	}

	// 如果都没有，返回默认语言
	return "en"
}
