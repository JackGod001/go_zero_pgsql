package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"go_zero_pgsql/common/globalkey"
	"net/http"
	"strings"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CasdoorJwtMiddleware struct {
	rsaPublicKey *rsa.PublicKey
	casdoorsdk   *casdoorsdk.Client
}

func NewCasdoorJwtMiddleware(casdoorClient *casdoorsdk.Client) (*CasdoorJwtMiddleware, error) {
	cerBy := []byte(casdoorClient.Certificate)
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(cerBy)
	if err != nil {
		return nil, err
	}
	return &CasdoorJwtMiddleware{
		rsaPublicKey: rsaPublicKey,
		casdoorsdk:   casdoorClient,
	}, nil
}
func (m *CasdoorJwtMiddleware) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return m.rsaPublicKey, nil
}
func (m *CasdoorJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) != 2 {
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}
		time.Sleep(500 * time.Millisecond)
		claims, err := m.casdoorsdk.ParseJwtToken(token[1])
		if err != nil {
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		//var (
		//	strToken string
		//	claims   jwt.RegisteredClaims
		//)
		//if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		//	strToken = strings.TrimPrefix(auth, "Bearer ")
		//}
		//if len(strToken) == 0 {
		//	httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
		//	return
		//}
		//
		//
		//// 我们明确设置为仅允许 RS256，并且还禁用
		//// 声明检查：RegisteredClaims 内部要求 'iat' 到
		//// 不晚于 'now'，但我们允许一点漂移。 todo 检查这里是否已经能验证时间
		//token, err := jwt.ParseWithClaims(strToken, &claims, m.keyFunc,
		//	jwt.WithValidMethods([]string{"RS256"}))
		//if err != nil {
		//	//未生效过期等验证失败
		//	httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
		//	return
		//}
		//
		ctx := r.Context()
		//// 设置用户ID的key
		ctx = context.WithValue(ctx, globalkey.SysJwtUserId, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
