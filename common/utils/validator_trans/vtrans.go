package validator_trans

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
)

func NewValidator() *validator.Validate {
	return validator.New()
}
func NewTranslator(validate *validator.Validate, lang string) (ut.Translator, error) {
	// 定义支持的语言
	matcher := language.NewMatcher([]language.Tag{
		language.Chinese,
		language.English,
	})
	// 解析语言标签
	tag, _ := language.MatchStrings(matcher, lang)

	var (
		translator           locales.Translator
		langString           string
		registerTranslations func(*validator.Validate, ut.Translator) error
	)

	// 根据匹配结果选择语言和翻译器
	switch tag {
	case language.Chinese:
		translator = zh.New()
		langString = "zh"
		registerTranslations = zhtranslations.RegisterDefaultTranslations
	default: // 包括 language.English 和其他未知语言
		translator = en.New()
		langString = "en"
		registerTranslations = entranslations.RegisterDefaultTranslations
	}

	// 创建翻译器
	trans, _ := ut.New(translator).GetTranslator(langString)

	// 注册默认翻译
	if err := registerTranslations(validate, trans); err != nil {
		return nil, err
	}

	return trans, nil
}
