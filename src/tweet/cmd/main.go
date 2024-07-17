package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	auth_grpc "github.com/fenek-dev/go-twitter/src/auth/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/middlewares"
	"github.com/fenek-dev/go-twitter/src/tweet/config"
	"github.com/fenek-dev/go-twitter/src/tweet/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/tweet/internal/services"
	"github.com/fenek-dev/go-twitter/src/tweet/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/tweet/internal/storage/redis"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

const (
	SERVICE_NAME = "tweet"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	tp := common.Init(ctx, SERVICE_NAME)
	defer tp.Shutdown(ctx)
	log := common.SetupLogger(cfg.Env)

	auth, err := auth_grpc.New(cfg.AuthUrl)
	if err != nil {
		panic("Could not connect to auth grpc server.")
	}
	defer auth.Close()

	auth_service := auth.NewService()

	tracer := otel.Tracer(SERVICE_NAME)

	postgres := pg.New(ctx, cfg.DbUrl, tracer)
	defer postgres.Close()

	rdb := redis.New(ctx, &cfg.Redis, tracer)
	defer rdb.Close()

	services := services.New(auth_service, postgres, rdb, tracer)

	handlers := handlers.New(services, log, tracer)

	auth_middleware := middlewares.NewAuthMiddleware(auth_service)

	r := gin.Default()
	r.Use(otelgin.Middleware(SERVICE_NAME))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	r.Use(c)

	v1 := r.Group("/api/v1")

	v1s := v1.Group("", auth_middleware.Handle())
	v1s.PUT("/tweet", handlers.CreateTweet)
	v1s.PATCH(" /tweet", handlers.UpdateTweet)
	v1s.DELETE("/tweet", handlers.DeleteTweet)

	go func() {
		r.Run(":" + cfg.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("Gracefully stopped")
}
