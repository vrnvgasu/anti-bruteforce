package postgres

import (
	"context"
	"fmt"
	"vrnvgasu/anti-bruteforce/internal/config"

	_ "github.com/jackc/pgx/stdlib" // postgresql provider
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type Service struct {
	db *sqlx.DB
}

func NewPostgresService() *Service {
	return &Service{}
}

func (s *Service) Connect(_ context.Context) error {
	db, err := sqlx.Open("pgx", config.Cfg.PSQL.DSN)
	if err != nil {
		return fmt.Errorf("open sql db fail: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("ping sql db fail: %w", err)
	}

	s.db = db

	return nil
}

func (s *Service) Close(_ context.Context) error {
	return s.db.Close()
}

func (s *Service) Migrate() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set dialect failed: %w", err)
	}

	if err := goose.Up(s.db.DB, config.Cfg.PSQL.Migration); err != nil {
		return fmt.Errorf("up migration failed: %w", err)
	}

	return nil
}
