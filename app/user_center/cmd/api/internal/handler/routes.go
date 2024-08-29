// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: GetUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user-center"),
		rest.WithTimeout(3000000*time.Millisecond),
	)
}
