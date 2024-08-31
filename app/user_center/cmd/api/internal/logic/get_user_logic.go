package logic

import (
	"context"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	userModel, err := l.svcCtx.UsersModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("UsersModel.FindOne", err.Error())
		return nil, err
	}
	return &types.GetUserResponse{
		Username:  userModel.Username,
		CreatedAt: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
