package handlers

import (
	"log/slog"

	"github.com/fenek-dev/go-twitter/src/tweet/internal/services"

	"go.opentelemetry.io/otel/trace"
)

type Handlers struct {
	service *services.Services

	log    *slog.Logger
	tracer trace.Tracer
}

func New(service *services.Services, log *slog.Logger, tracer trace.Tracer) *Handlers {
	return &Handlers{
		log:     log,
		service: service,
		tracer:  tracer,
	}
}
