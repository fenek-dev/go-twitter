package main

import "github.com/fenekdev/go-twitter/auth/config"

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
}
