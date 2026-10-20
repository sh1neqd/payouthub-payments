package main

import (
	"context"
	"payments/internal/app"
	"payments/internal/config"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := runApp(); err != nil {
		log.Printf("payments: %v", err)
		os.Exit(1)
	}
}

func runApp() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()

	if err != nil {
		return err
	}

	return app.Run(ctx, cfg)

}
