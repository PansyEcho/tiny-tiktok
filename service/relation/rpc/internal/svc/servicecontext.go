package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"tiny-tiktok/service/relation/model/model"
	"tiny-tiktok/service/relation/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	RedisCache *redis.Redis
	DB         *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		DB:         model.DB,
		Config:     c,
		RedisCache: c.RedisCacheConf.NewRedis(),
	}
}
