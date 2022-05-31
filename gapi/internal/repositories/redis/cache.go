package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) (*RedisRepo, error) {
	return &RedisRepo{
		client: client,
	}, nil
}

func (r RedisRepo) Get(key string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisRepo) Set(key string, value interface{}) error {
	//TODO implement me
	panic("implement me")
}
