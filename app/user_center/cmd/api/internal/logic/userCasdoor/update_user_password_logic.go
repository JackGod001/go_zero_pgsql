package userCasdoor

import (
	"context"
	"go_zero_pgsql/common/utils"

	"github.com/zeromicro/go-zero/core/errorx"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改密码
func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdatePasswordReq) error {

	userId := utils.GetCasdoorUserId(l.ctx)
	user, err := l.svcCtx.CasdoorClient.GetUserByUserId(userId)
	_, err = l.svcCtx.CasdoorClient.SetPassword(user.Owner, user.Name, req.OldPassword, req.NewPassword)
	if err != nil {
		return errorx.NewDefaultError(err.Error())
	}

	return nil
}
