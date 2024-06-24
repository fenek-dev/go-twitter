package main

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/cache/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/read-api/config"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	_ = common.SetupLogger(cfg.Env)

	storage := pg.New(ctx, cfg.DBUrl)
	handlers := handlers.New(storage)

}
