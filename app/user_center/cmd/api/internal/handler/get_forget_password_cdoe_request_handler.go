package handler

import (
	"go_zero_pgsql/app/user_center/cmd/api/internal/logic"
	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

// 忘记密码获取验证码
func GetForgetPasswordCdoeRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetForgetPasswordCdoeRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetForgetPasswordCdoeRequestLogic(r.Context(), svcCtx)
		err := l.GetForgetPasswordCdoeRequest(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
