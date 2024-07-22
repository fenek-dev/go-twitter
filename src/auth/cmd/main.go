package main

import (
	"context"
	"github.com/fenek-dev/go-twitter/src/auth/internal/http/middlewares"
	"os"
	"os/signal"
	"syscall"

	"github.com/fenek-dev/go-twitter/src/auth/config"
	"github.com/fenek-dev/go-twitter/src/auth/internal/grpc"
	"github.com/fenek-dev/go-twitter/src/auth/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/auth/internal/services"
	"github.com/fenek-dev/go-twitter/src/auth/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/auth/internal/storage/redis"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

const (
	SERVICE_NAME = "auth"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	trace := common.Init(ctx, SERVICE_NAME)
	defer trace.Shutdown(ctx)

	log := common.SetupLogger(cfg.Env)
	tracer := otel.Tracer(SERVICE_NAME)

	storage := pg.New(ctx, cfg.DBUrl, tracer)
	defer storage.Close()

	rdb := redis.New(ctx, &cfg.Redis, tracer)
	defer rdb.Close()

	service := services.New(storage, rdb, tracer, log, cfg.TokenTTL, cfg.Secret)

	grpcServer := grpc.New(log, service, tracer, cfg.GRPC.Port)

	go func() {
		grpcServer.MustRun()
	}()

	h := handlers.New(service, log, tracer)

	r := gin.Default()
	r.Use(otelgin.Middleware(SERVICE_NAME))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	r.Use(c)

	v1 := r.Group("/api/v1")
	v1.POST("/register", h.Register)
	v1.POST("/login", h.Login)

	v1s := v1.Group("", middlewares.Auth(service))
	v1s.GET("/me", h.Me)

	go func() {
		r.Run(":" + cfg.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	grpcServer.Stop()
	log.Info("Gracefully stopped")
}
