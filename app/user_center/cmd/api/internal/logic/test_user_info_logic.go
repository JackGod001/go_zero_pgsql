package logic

import (
	"context"
	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"
	"go_zero_pgsql/common/enum/errorcode"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息测试
func NewTestUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestUserInfoLogic {
	return &TestUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestUserInfoLogic) TestUserInfo(req *types.TestGetUserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//return &types.UserInfoResponse{
	//	Id:    1,
	//	Email: "test@qq.com",
	//}, nil
	return nil, errorx.NewCodeError(errorcode.InvalidArgument, "apiDesc.createConfiguration")
	//l.svcCtx.Trans.Trans(l.ctx, "apiDesc.createConfiguration")

}
