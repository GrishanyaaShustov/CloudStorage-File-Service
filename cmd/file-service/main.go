package main

import (
	"file-service/cmd/file-service/logger"
	"file-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("Hello world!")
}
