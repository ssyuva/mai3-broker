package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func NewRedisClient(url string) (*redis.Client, error) {
	opt, err := redis.ParseURL(url)

	if err != nil {
		return nil, fmt.Errorf("new redis client:parse client url fail:%w", err)
	}

	// Redis supports many configs.
	// You should change them by demand.
	opt.PoolSize = 10
	opt.MaxRetries = 2

	client := redis.NewClient(opt)
	return client, nil
}

var RedisClient *redis.Client

func Init(url string) error {
	var err error
	RedisClient, err = NewRedisClient(url)
	return err
}
