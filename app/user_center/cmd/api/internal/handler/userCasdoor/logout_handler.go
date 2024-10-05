package userCasdoor

import (
	"go_zero_pgsql/app/user_center/cmd/api/internal/logic/userCasdoor"
	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	xhttp "github.com/zeromicro/x/http"
)

// 退出
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userCasdoor.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
