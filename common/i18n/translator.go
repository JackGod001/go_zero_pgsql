package i18n

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"go_zero_pgsql/common/utils/errcode"
	"net/http"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/status"

	"io/fs"
	"path/filepath"
	"strings"

	"go_zero_pgsql/common/utils/parse"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

//go:embed locale/*.json
var LocaleFS embed.FS

// Translator is a struct storing translating data.
type Translator struct {
	bundle       *i18n.Bundle
	localizer    map[language.Tag]*i18n.Localizer
	supportLangs []language.Tag
}

// AddBundleFromEmbeddedFS adds new bundle into translator from embedded file system
func (l *Translator) AddBundleFromEmbeddedFS(file embed.FS, path string) error {
	if _, err := l.bundle.LoadMessageFileFS(file, path); err != nil {
		return err
	}
	return nil
}

// AddBundleFromFile adds new bundle into translator from file path.
func (l *Translator) AddBundleFromFile(path string) error {
	if _, err := l.bundle.LoadMessageFile(path); err != nil {
		return err
	}
	return nil
}

// AddLanguageSupport adds supports for new language
func (l *Translator) AddLanguageSupport(lang language.Tag) {
	l.supportLangs = append(l.supportLangs, lang)
	l.localizer[lang] = i18n.NewLocalizer(l.bundle, lang.String())
}

// Trans used to translate any i18n string.
func (l *Translator) Trans(ctx context.Context, msgId string) string {
	message, err := l.MatchLocalizer(ctx.Value("lang").(string)).LocalizeMessage(&i18n.Message{ID: msgId})
	if err != nil {
		return msgId
	}

	if message == "" {
		return msgId
	}

	return message
}
func (l *Translator) TransError(ctx context.Context, err error) error {
	// 从上下文中获取语言设置 todo 是否能获取到正确的语言设置
	lang, ok := ctx.Value("lang").(string)
	if !ok {
		// 如果无法获取语言设置，使用默认语言（假设为英语）
		lang = "en"
	}

	// 定义一个内部函数来处理本地化消息
	localizeMessage := func(messageID string) string {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: messageID})
		if e != nil || message == "" {
			return messageID // 如果本地化失败，返回原始消息
		}
		return message
	}

	// 处理不同类型的错误
	if errcode.IsGrpcError(err) {
		// 处理gRPC错误
		parts := strings.Split(err.Error(), "desc = ")
		if len(parts) > 1 {
			return status.Error(status.Code(err), localizeMessage(parts[1]))
		}
		// 如果错误信息不符合预期格式，则返回原始错误信息
		return status.Error(status.Code(err), localizeMessage(err.Error()))
	} else if codeErr, ok := err.(*errorx.CodeError); ok {
		return errorx.NewCodeError(codeErr.Code, localizeMessage(codeErr.Error()))
	} else if apiErr, ok := err.(*errorx.ApiError); ok {
		return errorx.NewApiError(apiErr.Code, localizeMessage(apiErr.Error()))
	} else {
		// 处理其他类型的错误
		return errorx.NewApiError(http.StatusInternalServerError, err.Error())
	}

}

// MatchLocalizer used to matcher the localizer in map
func (l *Translator) MatchLocalizer(lang string) *i18n.Localizer {
	tags := parse.ParseTags(lang)
	for _, v := range tags {
		if val, ok := l.localizer[v]; ok {
			return val
		}
	}

	return l.localizer[language.Chinese]
}

// NewTranslator returns a translator by I18n Conf.
// If Conf.Dir is empty, it will load paths in embedded FS.
// If Conf.Dir is not empty, it will load paths joined with Dir path.
// e.g. trans = i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)
func NewTranslator(conf Conf, efs embed.FS) *Translator {
	trans := &Translator{}
	trans.localizer = make(map[language.Tag]*i18n.Localizer)
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	trans.bundle = bundle

	var files []string
	if conf.Dir == "" {
		if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
			if d == nil {
				logx.Must(fmt.Errorf("wrong directory path: %s", conf.Dir))
			}
			if !d.IsDir() {
				files = append(files, path)
			}

			return err
		}); err != nil {
			logx.Must(fmt.Errorf("failed to get any files in dir: %s, error: %v", conf.Dir, err))
		}

		for _, v := range files {
			languageName := strings.TrimSuffix(filepath.Base(v), ".json")
			trans.AddLanguageSupport(parse.ParseTags(languageName)[0])
			err := trans.AddBundleFromEmbeddedFS(efs, v)
			if err != nil {
				logx.Must(fmt.Errorf("failed to load files from %s for i18n, please check the "+
					"configuration, error: %s", v, err.Error()))
			}
		}
	} else {
		if err := filepath.WalkDir(conf.Dir, func(path string, d fs.DirEntry, err error) error {
			if d == nil {
				logx.Must(fmt.Errorf("wrong directory path: %s", conf.Dir))
			}
			if !d.IsDir() {
				files = append(files, path)
			}

			return err
		}); err != nil {
			logx.Must(fmt.Errorf("failed to get any files in dir: %s, error: %v", conf.Dir, err))
		}

		for _, v := range files {
			languageName := strings.TrimSuffix(filepath.Base(v), ".json")
			trans.AddLanguageSupport(parse.ParseTags(languageName)[0])
			err := trans.AddBundleFromFile(v)
			if err != nil {
				logx.Must(fmt.Errorf("failed to load files from %s for i18n, please check the "+
					"configuration, error: %s", filepath.Join(conf.Dir, v), err.Error()))
			}
		}
	}

	return trans
}
