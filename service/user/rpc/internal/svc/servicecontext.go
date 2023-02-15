package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"tiny-tiktok/service/user/rpc/internal/config"
	"tiny-tiktok/service/user/rpc/internal/model"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	RedisCache *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	model.NewDB()
	return &ServiceContext{
		Config:     c,
		RedisCache: c.RedisCacheConf.NewRedis(),
		DB:         model.DB,
	}
}
