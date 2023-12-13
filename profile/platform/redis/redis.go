package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"profile/internal/cfg"
)

func StartRedis(config *cfg.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.RedisConfig.Host, config.RedisConfig.Port)
	})
	return client
}


