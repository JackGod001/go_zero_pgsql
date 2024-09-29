package logic

import (
	"context"
	"go_zero_pgsql/common/utils"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	userId := utils.GetUserId(l.ctx)

	userModel, err := l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("UsersModel.FindOne", err.Error())
		return nil, err
	}
	return &types.UserInfoResponse{
		Email: userModel.Email,
	}, nil

}
