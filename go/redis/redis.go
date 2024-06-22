package redis

import (
	"go-redis/config"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Connect(cfg *config.Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}
