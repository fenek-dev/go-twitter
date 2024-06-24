package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/storage"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	conn := pg.New(ctx, cfg.DBUrl)

	defer conn.Close(ctx)

	log := common.SetupLogger(cfg.Env)

	storage := storage.New(conn)

	sso, err := sso_grpc.NewSsoGrpcClient(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}

	services := services.New(sso)

	handlers := handlers.New(storage, services, log)

	http.HandleFunc("POST /api/v1/register", handlers.Register)
	http.HandleFunc("POST /api/v1/login", handlers.Login)

	http.HandleFunc("PUT /api/v1/tweet", handlers.CreateTweet)
	http.HandleFunc("PATCH /api/v1/tweet", handlers.UpdateTweet)
	http.HandleFunc("DELETE /api/v1/tweet", handlers.DeleteTweet)

	http.ListenAndServe(":"+cfg.Port, nil)
}
