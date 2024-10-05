package userCasdoor

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/errorx"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新token
func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.GetTokenByCodeResp, err error) {
	// 根据refreshToken刷新token
	token, err := l.svcCtx.CasdoorClient.RefreshOAuthToken(req.RefreshToken)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("刷新token失败", err.Error())
		return nil, errorx.NewApiErrorWithoutMsg(http.StatusInternalServerError)
	}
	return &types.GetTokenByCodeResp{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.Expiry.Unix(),
		TokenType:    token.TokenType,
	}, nil

}
