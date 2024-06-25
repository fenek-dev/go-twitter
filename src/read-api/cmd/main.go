package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/cache/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/read-api/config"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/handlers"
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

	handlers := handlers.New(cache)

	http.HandleFunc("GET /api/v1/tweet/{id}", handlers.FindTweetById)
	http.HandleFunc("GET /api/v1/user/{id}", handlers.FindUserById)

	go func() {
		http.ListenAndServe(":"+cfg.Port, nil)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	client.Close()

	log.Info("Gracefully stopped")
}
