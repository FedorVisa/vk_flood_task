package service

import (
	"context"
	"fmt"
	"log"
	"task/internal/config"
)

type FloodService struct {
	client *Redis
	config *config.Config
}

type FloodControl interface {
	Check(ctx context.Context, userID int64) (bool, error)
}

func NewService(client *Redis, config *config.Config) *FloodService {
	return &FloodService{
		client: client,
		config: config,
	}
}

func (s *FloodService) Check(ctx context.Context, userID int64) (bool, error) {
	key := fmt.Sprintf("UserID%d", userID)
	result, err := s.client.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		if err := s.client.client.Expire(ctx, key, s.config.K).Err(); err != nil {
			return false, err
		}
	}
	if result >= s.config.N {
		log.Println(fmt.Sprintf("ID: %d is flooding", userID))
		if err := s.client.client.Del(context.Background(), key).Err(); err != nil {
			return false, err
		}
		return false, err
	}
	return true, nil
}
