package main

import (
	"context"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/auth"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/tweets"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	conn := pg.New(ctx, cfg.DBUrl)

	defer conn.Close(ctx)

	tweet_repository := tweets.NewRepository(conn)

	log := common.SetupLogger(cfg.Env)

	sso, err := sso_grpc.NewSsoGrpcClient(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}

	auth_service := auth.NewService(sso)
	auth_controller := auth.NewController(log, auth_service)

	http.HandleFunc("POST /api/v1/register", auth_controller.Register)
	http.HandleFunc("POST /api/v1/login", auth_controller.Login)

	tweets_controller := tweets.NewController(tweet_repository)

	http.HandleFunc("PUT /api/v1/tweet", tweets_controller.Create)
	http.HandleFunc("PATCH /api/v1/tweet", tweets_controller.Update)
	http.HandleFunc("DELETE /api/v1/tweet", tweets_controller.Delete)

	http.ListenAndServe(":"+cfg.Port, nil)
}
