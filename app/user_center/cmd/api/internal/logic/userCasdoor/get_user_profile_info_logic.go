package userCasdoor

import (
	"context"
	"go_zero_pgsql/common/utils"

	"github.com/zeromicro/go-zero/core/errorx"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录信息，用于基础时直接获取的用户基础信息
func NewGetUserProfileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileInfoLogic {
	return &GetUserProfileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileInfoLogic) GetUserProfileInfo() (resp *types.UserProfileInfoResp, err error) {
	userId := utils.GetCasdoorUserId(l.ctx)
	//从casdoor 中获取资源
	user, err := l.svcCtx.CasdoorClient.GetUserByUserId(userId)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	return &types.UserProfileInfoResp{
		Id:       userId,
		Nickname: user.Name,
		//Email:    user.Email,
		Avatar: user.Avatar,
	}, nil
}
