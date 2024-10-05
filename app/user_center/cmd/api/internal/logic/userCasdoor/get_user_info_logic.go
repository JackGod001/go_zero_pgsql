package userCasdoor

import (
	"context"
	"go_zero_pgsql/common/utils"

	"github.com/zeromicro/go-zero/core/errorx"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详细信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoResp, err error) {
	user, err := utils.GetCasdoorUser(l.ctx, l.svcCtx.CasdoorClient)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.UserInfoResp{
		Username: user.Name,
		Avatar:   user.Avatar,
	}, nil
}
