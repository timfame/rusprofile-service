package main

import (
	"context"
	"github.com/timfame/rusprofile-service/internal/cache"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/rusprofile"
	"github.com/timfame/rusprofile-service/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	l := logger.New(logger.WithDebugLevel(), logger.WithServiceName("rusprofile"))

	cfg := config.New()

	err := cfg.Init(false)
	if err != nil {
		l.Fatal("Config initializing failed", zap.Error(err))
	}

	var c cache.Storage

	c, err = cache.NewRedis(context.Background(), cfg.Redis, cfg.Cache)
	if err != nil {
		l.Info("Cannot create Redis client", zap.Reflect("redis_config", cfg.Redis), zap.Error(err))
		c = cache.NewHashmap(cfg.Cache)
	}

	s := rusprofile.NewLogger(rusprofile.NewBase(cfg.Rusprofile, c), l)

	_, err = s.GetCompanyByINN(context.Background(), "7843007274")
	if err != nil {
		l.Fatal("error", zap.Error(err))
	}
}
