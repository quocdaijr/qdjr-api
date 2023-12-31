package initializers

import (
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisInitializer struct{}

var RedisClient *redis.Client

func (_ RedisInitializer) Redis() *redis.Client {
	return RedisClient
}

func (_ RedisInitializer) ConnectRedis() {
	url := baseHelper.GetEnv("REDIS_HOST", "127.0.0.1") + ":" + baseHelper.GetEnv("REDIS_PORT", "6379")
	db, _ := strconv.Atoi(baseHelper.GetEnv("REDIS_DB", "0"))
	username := baseHelper.GetEnv("REDIS_USERNAME", "")
	password := baseHelper.GetEnv("REDIS_PASSWORD", "")
	if username != "" && password != "" {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     url,
			Username: username,
			Password: password,
			DB:       db,
		})
	} else {
		RedisClient = redis.NewClient(&redis.Options{
			Addr: url,
			DB:   db,
		})
	}
}
