package handlers

import (
	"log/slog"

	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/otel/trace"
)

type Handlers struct {
	service *services.Services

	log    *slog.Logger
	tracer trace.Tracer
	Router *gin.Engine
}

func New(service *services.Services, log *slog.Logger, tracer trace.Tracer) *Handlers {
	return &Handlers{
		log:     log,
		service: service,
		tracer:  tracer,
		Router:  gin.Default(),
	}
}
