// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "tiny-tiktok/service/user/internal/handler/user"
	"tiny-tiktok/service/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)
}