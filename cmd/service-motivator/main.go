package main

import (
	"log/slog"

	"github.com/g-vinokurov/pyramidum-backend-service-motivator/internal/config"
	"github.com/g-vinokurov/pyramidum-backend-service-motivator/internal/env"
)

func main() {
	// init environment variables from .env file
	env.MustLoadEnv()

	cfg := config.MustLoadConfig()

	slog.Info("config loaded", slog.Any("config", cfg))
}
