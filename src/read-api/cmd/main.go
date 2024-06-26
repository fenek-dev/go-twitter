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
	"github.com/fenek-dev/go-twitter/src/read-api/config"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/handlers"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/rs/cors"
)

func main() {
	_ = context.Background()
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	client, err := client.New(cfg.CacheUrl)
	if err != nil {
		panic("Could not connect to cache grpc server.")
	}
	cache := client.NewService()

	sso, err := sso_grpc.New(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}
	sso_service := sso.NewService()

	handlers := handlers.New(cache)

	auth_middleware := middlewares.NewAuthMiddleware(sso_service)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/me", auth_middleware.Handle(handlers.Me))
	mux.HandleFunc("GET /api/v1/tweet/{id}", handlers.FindTweetById)
	mux.HandleFunc("GET /api/v1/user/{id}", handlers.FindUserById)
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

	client.Close()

	log.Info("Gracefully stopped")
}
