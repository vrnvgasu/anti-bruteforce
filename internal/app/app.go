package app

import (
	"context"
	"vrnvgasu/anti-bruteforce/internal/core/bucket"
)

type App struct {
	storage Storage
	bucket  Bucket
	l       Logger
}

type Logger interface {
	Warn(msg string)
	Info(msg string)
	Error(msg string)
	File(msg string)
}

type Storage interface {
	GetSet(ctx context.Context, set string) ([]string, error)
	AddToSet(ctx context.Context, set string, values ...string) error
	RemoveFromSet(ctx context.Context, set string, values ...string) error
	DeleteByKey(ctx context.Context, key string) error
}

type Bucket interface {
	Check(ctx context.Context, item string, entity bucket.EntityRate) (bool, error)
}

func New(logger Logger, storage Storage, bucket Bucket) *App {
	return &App{
		l:       logger,
		storage: storage,
		bucket:  bucket,
	}
}
