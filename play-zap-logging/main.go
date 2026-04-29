package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("info", zap.String("url", "some url"))
	logger.Warn("Warning", zap.Int64("int 640", 6400))
	logger.Fatal("failed to fethc url", zap.String("url", "some url"), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
}
