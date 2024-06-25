package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/cache/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/middlewares"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
)

func main() {
	_ = context.Background()
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	sso, err := sso_grpc.New(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}
	sso_service := sso.NewService()

	client, err := client.New(cfg.CacheUrl)
	if err != nil {
		panic("Could not connect to cache grpc server.")
	}
	cache := client.NewService()

	services := services.New(sso_service, cache)

	handlers := handlers.New(services, log)

	auth_middleware := middlewares.NewAuthMiddleware(sso_service)

	http.HandleFunc("POST /api/v1/register", handlers.Register)
	http.HandleFunc("POST /api/v1/login", handlers.Login)

	http.HandleFunc("PUT /api/v1/tweet", auth_middleware.Handle(handlers.CreateTweet))
	http.HandleFunc("PATCH /api/v1/tweet", auth_middleware.Handle(handlers.UpdateTweet))
	http.HandleFunc("DELETE /api/v1/tweet", auth_middleware.Handle(handlers.DeleteTweet))

	go func() {
		http.ListenAndServe(":"+cfg.Port, nil)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	sso.Close()
	client.Close()

	log.Info("Gracefully stopped")
}
