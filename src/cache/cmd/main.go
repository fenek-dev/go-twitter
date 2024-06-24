package main

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/cache/config"
	"github.com/fenek-dev/go-twitter/src/cache/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/redis"
	"github.com/fenek-dev/go-twitter/src/common"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	_ = common.SetupLogger(cfg.Env)

	storage := pg.New(ctx, cfg.DBUrl)
	redis := redis.New(cfg)

	handlers := handlers.New(storage, redis)

}
