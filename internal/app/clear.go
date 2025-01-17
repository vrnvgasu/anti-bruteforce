package app

import (
	"context"
	"fmt"
)

type ClearRequest struct {
	Login string `json:"login"`
	IP    string `json:"ip"`
}

func (a *App) Clear(ctx context.Context, req ClearRequest) error {
	if err := a.storage.DeleteByKey(ctx, req.IP); err != nil {
		return fmt.Errorf("app Clear ip DeleteByKey: %w", err)
	}
	if err := a.storage.DeleteByKey(ctx, req.Login); err != nil {
		return fmt.Errorf("app Clear login DeleteByKey: %w", err)
	}

	return nil
}
