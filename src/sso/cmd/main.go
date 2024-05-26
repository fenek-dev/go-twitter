package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	app "github.com/fenek-dev/go-twitter/src/sso/app/grpc"
	"github.com/fenek-dev/go-twitter/src/sso/config"
	"github.com/fenek-dev/go-twitter/src/sso/internal/services/auth"
	"github.com/fenek-dev/go-twitter/src/sso/storage/pg"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	storage := pg.New(ctx, cfg.DBUrl)

	log := config.SetupLogger(cfg.Env)

	auth_service := auth.New(log, storage, cfg.TokenTTL, cfg.Secret)

	grpc_server := app.New(log, auth_service, cfg.GRPC.Port)

	go func() {
		grpc_server.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	grpc_server.Stop()
	storage.Stop(ctx)
	log.Info("Gracefully stopped")
}
