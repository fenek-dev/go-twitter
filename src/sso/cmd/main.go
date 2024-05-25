package main

import (
	"github.com/fenek-dev/go-twitter/src/sso/config"
)

func main() {
	cfg := config.MustLoad()

	log := config.SetupLogger(cfg.Env)

	log.Debug("sso")
}
