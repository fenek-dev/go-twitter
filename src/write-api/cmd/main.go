package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/cache/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
)

func main() {
	_ = context.Background()
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	sso, err := sso_grpc.NewSsoGrpcClient(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}
	client, err := client.New(cfg.CacheUrl)
	if err != nil {
		panic("Could not connect to cache grpc server.")
	}
	cache := client.NewService()

	services := services.New(sso, cache)

	handlers := handlers.New(services, log)

	http.HandleFunc("POST /api/v1/register", handlers.Register)
	http.HandleFunc("POST /api/v1/login", handlers.Login)

	http.HandleFunc("PUT /api/v1/tweet", handlers.CreateTweet)
	http.HandleFunc("PATCH /api/v1/tweet", handlers.UpdateTweet)
	http.HandleFunc("DELETE /api/v1/tweet", handlers.DeleteTweet)

	http.ListenAndServe(":"+cfg.Port, nil)
}
