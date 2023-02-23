package svc

import (
	"tiny-tiktok/common/kqueue"
	"tiny-tiktok/service/mqueue/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	KqueueClient kqueue.KqueueClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		KqueueClient: kqueue.NewKqueueSvcClient(c.KqServerConf.Brokers),
	}
}
