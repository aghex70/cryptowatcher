package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RedisRepo struct {
	client *redis.Client
	logger *zap.Logger
	ctx    context.Context
}

func NewRedisRepo(client *redis.Client) (*RedisRepo, error) {
	return &RedisRepo{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (r RedisRepo) Get(key string) (interface{}, error) {
	value, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			r.logger.Warn(fmt.Sprintf("Key %s does not exist", key))
		}
		return nil, err
	}
	return value, nil
}

func (r RedisRepo) Set(key string, value interface{}) error {
	err := r.client.Set(r.ctx, key, value, 30).Err()
	if err != nil {
		return err
	}
	return nil
}
