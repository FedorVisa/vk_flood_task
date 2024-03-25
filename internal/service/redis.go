package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Redis struct {
	client *redis.Client
}

func MakeRedisService(addr string) *Redis {

	redisServ := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if err := redisServ.FlushAll(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to flush Redis data: %v", err)
	}
	return &Redis{client: redisServ}

}
