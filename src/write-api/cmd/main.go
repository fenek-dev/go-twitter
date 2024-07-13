package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/cache/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/middlewares"
	sso_grpc "github.com/fenek-dev/go-twitter/src/sso/pkg/client"
	"github.com/fenek-dev/go-twitter/src/write-api/config"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
	cors "github.com/rs/cors/wrapper/gin"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

const (
	SERVICE_NAME = "write-api"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	tp := common.Init(ctx, SERVICE_NAME)
	defer tp.Shutdown(ctx)
	log := common.SetupLogger(cfg.Env)

	sso, err := sso_grpc.New(cfg.SsoUrl)
	if err != nil {
		panic("Could not connect to sso grpc server.")
	}
	sso_service := sso.NewService()

	client, err := client.New(cfg.CacheUrl)
	if err != nil {
		panic("Could not connect to cache grpc server.")
	}
	cache := client.NewService()

	tracer := otel.Tracer(SERVICE_NAME)

	services := services.New(sso_service, cache, tracer)

	handlers := handlers.New(services, log, tracer)

	auth_middleware := middlewares.NewAuthMiddleware(sso_service)

	r := handlers.Router
	r.Use(otelgin.Middleware(SERVICE_NAME))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	r.Use(c)

	v1 := r.Group("/api/v1")
	v1.POST("/register", handlers.Register)
	v1.POST("/login", handlers.Login)

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

	sso.Close()
	client.Close()

	log.Info("Gracefully stopped")
}
