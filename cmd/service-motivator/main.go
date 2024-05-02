package main

import (
	"github.com/g-vinokurov/pyramidum-backend-service-motivator/internal/config"
	"github.com/g-vinokurov/pyramidum-backend-service-motivator/internal/env"
	"log/slog"
	"os"
)

func main() {
	// init environment variables from .env file
	env.MustLoadEnv()

	configPath := os.Getenv("CONFIG_PATH")
	cfg := config.MustLoadConfig(configPath)

	slog.Info("config loaded", slog.Any("config", cfg))
}
