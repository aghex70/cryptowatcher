package persistence

import (
	"fmt"
	"cwapi/config"
	"cwapi/internal/logger"
	"github.com/go-redis/redis/v8"
)

func NewRedisCache() (*redis.Client, error) {
	address := fmt.Sprintf("%s:%d", config.C.Cache.Host, config.C.Cache.Port)
	logger.ZapLogger.Info("Connecting to Redis")
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.C.Cache.Password,
		DB:       config.C.Cache.DB,
	})

	return rdb, nil
}
