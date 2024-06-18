package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/auth"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	storage := pg.New(ctx, cfg.DBUrl)

	defer storage.Close(ctx)

	log := common.SetupLogger(cfg.Env)

	auth_service := auth.NewService()
	auth_controller := auth.NewController(log, auth_service)

	http.HandleFunc("/register", auth_controller.Register)
}
