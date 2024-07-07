package redis

import (
	setting "GoBBS/settings"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(config *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})
	if _, err = rdb.Ping(context.Background()).Result(); err != nil {
		zap.L().Error("fail to connect to redis", zap.Error(err))
		return
	}
	return
}
