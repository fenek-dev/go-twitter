package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/fenek-dev/go-twitter/src/read-api/config"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/auth"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/tweets"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/user"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	conn := pg.New(ctx, cfg.DBUrl)

	defer conn.Close(ctx)

	tweet_repository := tweets.NewRepository(conn)
	user_repository := user.NewRepository(conn)

	log := common.SetupLogger(cfg.Env)

	sso, err := sso_grpc.NewSsoGrpcClient(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}

	auth_service := auth.NewService(sso)
	_ = auth.NewController(log, auth_service)

	tweets_controller := tweets.NewController(tweet_repository)
	user_controller := user.NewController(user_repository)

	http.HandleFunc("GET /api/v1/tweet/{id}", tweets_controller.FindById)
	http.HandleFunc("GET /api/v1/user/{id}", user_controller.FindById)

	http.ListenAndServe(":"+cfg.Port, nil)
}
