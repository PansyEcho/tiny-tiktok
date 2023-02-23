package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"tiny-tiktok/service/relation/internal/config"
	"tiny-tiktok/service/relation/rpc/relationservice"
)

type ServiceContext struct {
	Config      config.Config
	RedisCache  *redis.Redis
	RelationRpc relationservice.RelationService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisCache:  c.RedisCacheConf.NewRedis(),
		RelationRpc: relationservice.NewRelationService(zrpc.MustNewClient(c.RelationRpcConf)),
	}
}
