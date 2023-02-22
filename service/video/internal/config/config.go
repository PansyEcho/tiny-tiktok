package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RedisCacheConf redis.RedisConf

	COSConf struct {
		SecretId    string
		SecretKey   string
		MachineId   uint16
		VideoBucket string
		CoverBucket string
	}
}
