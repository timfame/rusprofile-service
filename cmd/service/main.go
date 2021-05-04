package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/timfame/rusprofile-service/internal/cache"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/delivery/grpc"
	"github.com/timfame/rusprofile-service/internal/rusprofile"
	"github.com/timfame/rusprofile-service/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	grpcServ := grpc.NewServer(cfg.Grpc, s, l)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return errors.WithMessage(grpcServ.Run(), "Grpc server")
	})
	g.Go(func() error {
		return errors.WithMessage(grpcServ.RunGateway(ctx), "Grpc gateway server")
	})

	select {
	case sig := <-interrupt:
		l.Info("Exiting program", zap.String("reason", sig.String()))
	case <-ctx.Done():
		break
	}
	l.Info("Shutdown signal received")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer shutdownCancel()

	if err := grpcServ.GracefulStop(shutdownCtx); err != nil {
		l.Info("Grpc gateway server graceful stop failed", zap.Error(err))
	}

	if err := g.Wait(); err != nil {
		l.Info("Serving one of the servers failed", zap.Error(err))
	}
}
