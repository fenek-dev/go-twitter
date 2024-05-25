package main

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/sso/config"
	"github.com/fenek-dev/go-twitter/src/sso/storage/pg"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	storage := pg.New(ctx, cfg.DBUrl)

	log := config.SetupLogger(cfg.Env)

	log.Debug("sso")
}
