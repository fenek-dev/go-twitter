package main

import (
	"context"
	"net/http"
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
	"github.com/rs/cors"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	tracer := common.Init(ctx, "write-api")

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

	services := services.New(sso_service, cache, tracer)

	handlers := handlers.New(services, log, tracer)

	auth_middleware := middlewares.NewAuthMiddleware(sso_service)

	mux := http.NewServeMux()
	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	handleFunc("POST /api/v1/register", handlers.Register)
	handleFunc("POST /api/v1/login", handlers.Login)

	handleFunc("PUT /api/v1/tweet", auth_middleware.Handle(handlers.CreateTweet))
	handleFunc("PATCH /api/v1/tweet", auth_middleware.Handle(handlers.UpdateTweet))
	handleFunc("DELETE /api/v1/tweet", auth_middleware.Handle(handlers.DeleteTweet))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	corsHandler := c.Handler(mux)

	handler := otelhttp.NewHandler(corsHandler, "/")

	go func() {
		http.ListenAndServe(":"+cfg.Port, handler)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	sso.Close()
	client.Close()

	log.Info("Gracefully stopped")
}
