package redis

import (
	"context"
	"fmt"
	"vrnvgasu/anti-bruteforce/internal/config"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	client redis.UniversalClient
}

func NewRedisService() *Service {
	return &Service{}
}

func (s *Service) Connect(ctx context.Context) error {
	s.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
		Password: config.Cfg.Password,
		Username: config.Cfg.Username,
		DB:       config.Cfg.DB,
	})
	_, err := s.client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis service: %w", err)
	}

	return nil
}

func (s *Service) Close(ctx context.Context) error {
	s.client.Shutdown(ctx)

	return nil
}
