package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/auth"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	storage := pg.New(ctx, cfg.DBUrl)

	defer storage.Close(ctx)

	log := common.SetupLogger(cfg.Env)

	sso, err := sso_grpc.NewSsoGrpcClient(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}

	auth_service := auth.NewService(sso)
	auth_controller := auth.NewController(log, auth_service)

	http.HandleFunc("/api/v1/register", auth_controller.Register)
	http.HandleFunc("/api/v1/login", auth_controller.Login)
}
