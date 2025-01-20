package storage

import (
	"context"
	"fmt"
	"time"
)

type KeyValueStorage interface {
	IncrCounterByKey(ctx context.Context, key string) (int64, error)
	ExpireKey(ctx context.Context, key string, duration time.Duration) error

	DeleteByKey(ctx context.Context, key string) error

	GetSet(ctx context.Context, set string) ([]string, error)
	AddToSet(ctx context.Context, set string, values ...string) error
	RemoveFromSet(ctx context.Context, set string, values ...string) error

	Connect(_ context.Context) error
	Close(_ context.Context) error
}

type Storage struct {
	KeyValueStorage
}

func NewStorage(kv KeyValueStorage) *Storage {
	return &Storage{KeyValueStorage: kv}
}

func (s *Storage) Start(ctx context.Context) error {
	if err := s.KeyValueStorage.Connect(ctx); err != nil {
		return fmt.Errorf("storage: failed to connect to key value storage: %w", err)
	}

	return nil
}

func (s *Storage) ShutDown() error {
	if s.KeyValueStorage != nil {
		if err := s.KeyValueStorage.Close(context.Background()); err != nil {
			return fmt.Errorf("storage: failed to close key value storage: %w", err)
		}
	}

	return nil
}
