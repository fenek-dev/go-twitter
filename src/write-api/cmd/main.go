package main

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	storage := pg.New(ctx, cfg.DBUrl)

	defer storage.Close(ctx)

	log := common.SetupLogger(cfg.Env)
}
