package app

import (
	"context"
	"log"
	httpserver "payments/internal/adapters/http"
	"payments/internal/adapters/postgres"
	"payments/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	httpServer := httpserver.NewServer(cfg)

	if err := httpServer.Run(ctx); err != nil {
		return err
	}

	log.Println("Payments server stopped")

	return nil
}
