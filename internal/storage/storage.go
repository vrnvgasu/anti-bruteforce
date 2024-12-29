package storage

import (
	"context"
	"fmt"
)

type DBStorage interface {
	Connect(_ context.Context) error
	Close(_ context.Context) error
	Migrate() error
}

type KeyValueStorage interface {
	Connect(_ context.Context) error
	Close(_ context.Context) error
}

type Storage struct {
	db DBStorage
	kv KeyValueStorage
}

func NewStorage(db DBStorage, kv KeyValueStorage) *Storage {
	return &Storage{db: db, kv: kv}
}

func (s *Storage) Start(ctx context.Context) error {
	if err := s.db.Connect(ctx); err != nil {
		return fmt.Errorf("storage: failed to connect to db storage: %w", err)
	}
	if err := s.kv.Connect(ctx); err != nil {
		return fmt.Errorf("storage: failed to connect to key value storage: %w", err)
	}

	return nil
}

func (s *Storage) ShutDown() error {
	if s.db != nil {
		if err := s.db.Close(context.Background()); err != nil {
			return fmt.Errorf("storage: failed to close db storage: %w", err)
		}
	}
	if s.kv != nil {
		if err := s.kv.Close(context.Background()); err != nil {
			return fmt.Errorf("storage: failed to close key value storage: %w", err)
		}
	}

	return nil
}
