package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/cache/config"
	"github.com/fenek-dev/go-twitter/src/cache/internal/grpc"
	"github.com/fenek-dev/go-twitter/src/cache/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/redis"
	"github.com/fenek-dev/go-twitter/src/common"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	storage := pg.New(ctx, cfg.DBUrl)
	redis := redis.New(&cfg.Redis)

	handlers := handlers.New(storage, redis)

	grpc_server := grpc.New(log, handlers, cfg.GRPC.Port)

	go grpc_server.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	grpc_server.Stop()
	log.Info("Gracefully stopped")
}
