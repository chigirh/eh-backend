package redis

import (
	"eh-backend-api/conf/config"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// go-redis
// SEE:https://github.com/go-redis/redis
func NewClient() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Config.Redis.RedisHost, config.Config.Redis.RedisPort),
		Password: "",
		DB:       0,
	})
	return cli
}
