package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	app "github.com/fenek-dev/go-twitter/src/sso/app/grpc"
	"github.com/fenek-dev/go-twitter/src/sso/config"
	user_domain "github.com/fenek-dev/go-twitter/src/sso/internal/domains/user"
	"github.com/fenek-dev/go-twitter/src/sso/internal/services/auth"
	"go.opentelemetry.io/otel"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	storage := pg.New(ctx, cfg.DBUrl)

	trace := common.Init(ctx, "auth")
	defer trace.Shutdown(ctx)

	defer storage.Close(ctx)

	log := common.SetupLogger(cfg.Env)

	tracer := otel.Tracer("auth")

	user_repository := user_domain.NewRepository(storage, tracer)

	auth_service := auth.New(tracer, log, user_repository, cfg.TokenTTL, cfg.Secret)

	grpc_server := app.New(log, auth_service, tracer, cfg.GRPC.Port)

	go func() {
		grpc_server.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	grpc_server.Stop()
	log.Info("Gracefully stopped")
}
