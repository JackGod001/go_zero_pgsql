package jwt

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/zeromicro/go-zero/core/timex"
)

const claimHistoryResetDuration = time.Hour * 24

type (
	// ParseOption defines the method to customize a TokenParser.
	ParseOption func(parser *TokenParser)

	// A TokenParser is used to parse tokens.
	// TokenParser 结构体用于解析令牌。
	// 它维护了一个令牌历史记录的同步Map，用于跟踪和验证令牌。
	// 两个时间持续期字段用于管理令牌的重置和过期策略。
	TokenParser struct {
		resetTime     time.Duration // 令牌重置时间间隔，决定何时清除旧的令牌记录。
		resetDuration time.Duration // 令牌的有效持续时间，标识令牌在多长时间内被认为是有效的。
		history       sync.Map      // 一个同步Map，用于存储和查找令牌，确保在并发环境下的线程安全。
	}
)

// NewTokenParser returns a TokenParser.
func NewTokenParser(opts ...ParseOption) *TokenParser {
	// 创建一个TokenParser实例，并设置默认的resetTime和resetDuration。
	parser := &TokenParser{
		resetTime: timex.Now(),
		// 令牌的有效持续时间，标识令牌在多长时间内被认为是有效的。
		resetDuration: claimHistoryResetDuration,
	}
	// 应用用户自定义的ParseOption，允许用户自定义TokenParser的设置。
	for _, opt := range opts {
		// 调用ParseOption函数，传入TokenParser实例，允许用户自定义TokenParser的设置。
		opt(parser)
	}

	return parser
}

// ParseToken parses token from given r, with passed in secret and prevSecret.
// ParseToken方法用于从给定的请求中解析令牌，并使用指定的密钥和前一个密钥。
func (tp *TokenParser) ParseToken(r *http.Request, secret, prevSecret string) (*jwt.Token, error) {
	// 初始化一个令牌变量和错误变量。
	var token *jwt.Token
	var err error
	// 如果prevSecret不为空，则尝试使用两个密钥进行解析。
	if len(prevSecret) > 0 {
		// 获取两个密钥的计数，并比较它们。
		count := tp.loadCount(secret)
		prevCount := tp.loadCount(prevSecret)

		var first, second string

		// 如果第一个密钥的计数大于第二个密钥的计数，则将第一个密钥和第二个密钥交换。
		if count > prevCount {
			first = secret
			second = prevSecret
		} else {
			first = prevSecret
			second = secret
		}

		token, err = tp.doParseToken(r, first)
		if err != nil {
			token, err = tp.doParseToken(r, second)
			if err != nil {
				return nil, err
			}

			tp.incrementCount(second)
		} else {
			tp.incrementCount(first)
		}
	} else {
		token, err = tp.doParseToken(r, secret)
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

// 这是一个私有方法，用于从请求中解析令牌，并使用指定的密钥进行解析。
func (tp *TokenParser) doParseToken(r *http.Request, secret string) (*jwt.Token, error) {
	return request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		}, request.WithParser(newParser()))
}

func (tp *TokenParser) incrementCount(secret string) {
	now := timex.Now()
	if tp.resetTime+tp.resetDuration < now {
		tp.history.Range(func(key, value any) bool {
			tp.history.Delete(key)
			return true
		})
	}

	value, ok := tp.history.Load(secret)
	if ok {
		atomic.AddUint64(value.(*uint64), 1)
	} else {
		var count uint64 = 1
		tp.history.Store(secret, &count)
	}
}

func (tp *TokenParser) loadCount(secret string) uint64 {
	value, ok := tp.history.Load(secret)
	if ok {
		return *value.(*uint64)
	}

	return 0
}

// WithResetDuration returns a func to customize a TokenParser with reset duration.
func WithResetDuration(duration time.Duration) ParseOption {
	return func(parser *TokenParser) {
		parser.resetDuration = duration
	}
}

func newParser() *jwt.Parser {
	return jwt.NewParser(jwt.WithJSONNumber())
}
