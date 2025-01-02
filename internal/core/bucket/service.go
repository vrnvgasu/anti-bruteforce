package bucket

import (
	"context"
	"fmt"
	"time"
	"vrnvgasu/anti-bruteforce/internal/config"
)

//go:generate go run github.com/vektra/mockery/v2@v2.50.1 --name=Storage --with-expecter
type Storage interface {
	IncrCounterByKey(ctx context.Context, key string) (int64, error)
	ExpireKey(ctx context.Context, key string, duration time.Duration) error
}

type EntityRate int

const (
	Login EntityRate = iota
	Password
	IP
)

type Service struct {
	loginRate    int64
	passwordRate int64
	ipRate       int64
	timeLimit    time.Duration
	st           Storage
}

func NewBucket(st Storage) *Service {
	return &Service{
		loginRate:    config.Cfg.Rate.Login,
		passwordRate: config.Cfg.Rate.Password,
		ipRate:       config.Cfg.Rate.IP,
		timeLimit:    time.Duration(config.Cfg.Rate.TimeLimit) * time.Second,
		st:           st,
	}
}

func (s *Service) Check(ctx context.Context, item string, entity EntityRate) (bool, error) {
	switch entity {
	case Login:
		return s.checkByStorage(ctx, item, s.loginRate)
	case Password:
		return s.checkByStorage(ctx, item, s.passwordRate)
	case IP:
		return s.checkByStorage(ctx, item, s.ipRate)
	default:
		return false, fmt.Errorf("unknown entity")
	}
}

func (s *Service) checkByStorage(ctx context.Context, item string, rate int64) (bool, error) {
	counter, err := s.st.IncrCounterByKey(ctx, item)
	if err != nil {
		return false, fmt.Errorf("bucket checkByStorage checkByStorage: %w", err)
	}

	if counter == 1 {
		if err = s.st.ExpireKey(ctx, item, s.timeLimit); err != nil {
			return false, fmt.Errorf("bucket checkByStorage ExpireKey: %w", err)
		}
	}

	if counter < rate {
		return true, nil
	}

	return false, nil
}
