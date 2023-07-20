package redis

import (
	"context"
	"elasticSearch/internal/configs"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *configs.Configs) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Host + ":" + cfg.Cache.Port,
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.DB,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
