package redis

import (
	"context"
	"errors"
	"fmt"
	"time"
	"vrnvgasu/anti-bruteforce/internal/config"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	client redis.UniversalClient
}

func NewRedisService() *Service {
	return &Service{}
}

func (s *Service) IncrCounterByKey(ctx context.Context, key string) (int64, error) {
	counter, err := s.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis IncrCounterByKey: %w", err)
	}

	return counter, nil
}

func (s *Service) ExpireKey(ctx context.Context, key string, duration time.Duration) error {
	_, err := s.client.Expire(ctx, key, duration).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis ExpireKey: %w", err)
	}

	return nil
}

func (s *Service) DeleteByKey(ctx context.Context, key string) error {
	_, err := s.client.Del(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis DeleteByKey: %w", err)
	}

	return nil
}

func (s *Service) GetSet(ctx context.Context, set string) ([]string, error) {
	res, err := s.client.SMembers(ctx, set).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("redis GetSet: %w", err)
	}

	return res, nil
}

func (s *Service) AddToSet(ctx context.Context, set string, values ...string) error {
	_, err := s.client.SAdd(ctx, set, values).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis AddToSet: %w", err)
	}

	return nil
}

func (s *Service) RemoveFromSet(ctx context.Context, set string, values ...string) error {
	_, err := s.client.SRem(ctx, set, values).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis RemoveFromSet: %w", err)
	}

	return nil
}

func (s *Service) Connect(ctx context.Context) error {
	s.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
		Password: config.Cfg.Redis.Password,
		Username: config.Cfg.Redis.Username,
		DB:       config.Cfg.Redis.DB,
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
