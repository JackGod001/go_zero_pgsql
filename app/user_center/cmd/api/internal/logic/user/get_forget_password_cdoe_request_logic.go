package user

import (
	"context"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetForgetPasswordCdoeRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 忘记密码获取验证码
func NewGetForgetPasswordCdoeRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetForgetPasswordCdoeRequestLogic {
	return &GetForgetPasswordCdoeRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetForgetPasswordCdoeRequestLogic) GetForgetPasswordCdoeRequest(req *types.GetForgetPasswordCdoeRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
