package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/fenek-dev/go-twitter/src/read-api/config"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/storage"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	conn := pg.New(ctx, cfg.DBUrl)

	defer conn.Close(ctx)

	_ = common.SetupLogger(cfg.Env)

	storage := storage.New(conn)

	handlers := handlers.New(storage)

	http.HandleFunc("GET /api/v1/tweet/{id}", handlers.FindTweetById)
	http.HandleFunc("GET /api/v1/user/{id}", handlers.FindUserById)

	http.ListenAndServe(":"+cfg.Port, nil)
}
