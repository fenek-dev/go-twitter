package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/cache/config"
	"github.com/fenek-dev/go-twitter/src/cache/internal/grpc"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/redis"
	"github.com/fenek-dev/go-twitter/src/common"
	"go.opentelemetry.io/otel"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	tp := common.Init(ctx, "cache")
	defer tp.Shutdown(ctx)

	log := common.SetupLogger(cfg.Env)

	tracer := otel.Tracer("cache")

	storage := pg.New(ctx, cfg.DBUrl, tracer)
	redis := redis.New(&cfg.Redis)

	grpc_server := grpc.New(log, storage, redis, cfg.GRPC.Port)

	go grpc_server.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	grpc_server.Stop()
	log.Info("Gracefully stopped")
}
