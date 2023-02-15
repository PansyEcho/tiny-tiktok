package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tiny-tiktok/service/interaction/internal/config"
	"tiny-tiktok/service/user/rpc/userinfoservice"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userinfoservice.UserInfoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userinfoservice.NewUserInfoService(zrpc.MustNewClient(c.UserRpc)),
	}
}
