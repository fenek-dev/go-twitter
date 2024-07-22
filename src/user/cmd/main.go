package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	authgrpc "github.com/fenek-dev/go-twitter/src/auth/pkg/client"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/user/config"
	"github.com/fenek-dev/go-twitter/src/user/internal/handlers"
	"github.com/fenek-dev/go-twitter/src/user/internal/services"
	"github.com/fenek-dev/go-twitter/src/user/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/user/internal/storage/redis"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

const (
	SERVICE_NAME = "user"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	tp := common.Init(ctx, SERVICE_NAME)
	defer tp.Shutdown(ctx)
	log := common.SetupLogger(cfg.Env)

	auth, err := authgrpc.New(cfg.AuthUrl)
	if err != nil {
		panic("Could not connect to auth grpc server.")
	}
	defer auth.Close()

	authService := auth.NewService()

	tracer := otel.Tracer(SERVICE_NAME)

	postgres := pg.New(ctx, cfg.DbUrl, tracer)
	defer postgres.Close()

	rdb := redis.New(ctx, &cfg.Redis, tracer)
	defer rdb.Close()

	s := services.New(authService, postgres, rdb, tracer)

	h := handlers.New(s, log, tracer)

	r := gin.Default()
	r.Use(otelgin.Middleware(SERVICE_NAME))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	r.Use(c)

	v1 := r.Group("/api/v1")
	v1.GET("/user/:id", h.FindUserById)

	go func() {
		r.Run(":" + cfg.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("Gracefully stopped")
}
