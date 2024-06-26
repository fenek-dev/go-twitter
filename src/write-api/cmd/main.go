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
	"github.com/rs/cors"
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

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/register", handlers.Register)
	mux.HandleFunc("POST /api/v1/login", handlers.Login)

	mux.HandleFunc("PUT /api/v1/tweet", auth_middleware.Handle(handlers.CreateTweet))
	mux.HandleFunc("PATCH /api/v1/tweet", auth_middleware.Handle(handlers.UpdateTweet))
	mux.HandleFunc("DELETE /api/v1/tweet", auth_middleware.Handle(handlers.DeleteTweet))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	go func() {
		http.ListenAndServe(":"+cfg.Port, handler)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	sso.Close()
	client.Close()

	log.Info("Gracefully stopped")
}
