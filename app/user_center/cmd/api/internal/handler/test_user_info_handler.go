package handler

import (
	"go_zero_pgsql/app/user_center/cmd/api/internal/logic"
	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 获取用户信息测试
func TestUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestGetUserInfoRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTestUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.TestUserInfo(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
