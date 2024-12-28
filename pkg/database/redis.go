package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/ryvasa/go-super-farmer/pkg/env"
)

func NewRedisClient(env *env.Env) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     env.Redis.Host + ":" + env.Redis.Port, // misalnya "localhost:6379"
		Password: env.Redis.Password,                    // password kosong
		DB:       0,                                     // default DB
	})
}
