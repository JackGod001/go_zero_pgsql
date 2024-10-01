package middleware

import (
	"context"
	"go_zero_pgsql/common/utils/request"
	"net/http"
)

type I18nMiddleware struct {
}

func NewI18nMiddleware() *I18nMiddleware {
	return &I18nMiddleware{}
}

type contextKey string

const langContextKey contextKey = "lang"

// SetLanguageToContext 将语言设置存储到上下文中
func SetLanguageToContext(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, langContextKey, lang)
}

// GetLanguageFromContext 从上下文中获取语言设置
func GetLanguageFromContext(ctx context.Context) string {
	lang, ok := ctx.Value(langContextKey).(string)
	if !ok {
		return "en" // 默认语言
	}
	return lang
}

func (m *I18nMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := request.GetRequestLang(r)

		// 将语言设置存储在请求上下文中
		ctx := SetLanguageToContext(r.Context(), lang)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
