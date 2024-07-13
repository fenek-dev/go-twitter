package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/cache/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/middlewares"
	"github.com/fenek-dev/go-twitter/src/read-api/config"
	"github.com/fenek-dev/go-twitter/src/read-api/internal/handlers"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

const (
	SERVICE_NAME = "read-api"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	tp := common.Init(ctx, SERVICE_NAME)
	defer tp.Shutdown(ctx)
	log := common.SetupLogger(cfg.Env)

	client, err := client.New(cfg.CacheUrl)
	if err != nil {
		panic("Could not connect to cache grpc server.")
	}
	cache := client.NewService()

	sso, err := sso_grpc.New(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}
	sso_service := sso.NewService()

	tracer := otel.Tracer(SERVICE_NAME)

	handlers := handlers.New(cache, tracer)

	r := gin.Default()
	r.Use(otelgin.Middleware(SERVICE_NAME))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	r.Use(c)

	v1 := r.Group("/api/v1")
	v1.GET("/tweet/:id", handlers.FindTweetById)
	v1.GET("/user/:id", handlers.FindUserById)

	auth_middleware := middlewares.NewAuthMiddleware(sso_service)
	v1s := v1.Group("", auth_middleware.Handle())
	v1s.GET("/me", handlers.Me)

	go func() {
		r.Run(":" + cfg.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	client.Close()

	log.Info("Gracefully stopped")
}
