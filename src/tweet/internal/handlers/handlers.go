package handlers

import (
	"context"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	FindTweetById(ctx context.Context, id string) (*models.Tweet, error)
	CreateTweet(ctx context.Context, username, content string) (*models.Tweet, error)
	UpdateTweet(ctx context.Context, id, content string) (*models.Tweet, error)
	DeleteTweet(ctx context.Context, id string) (string, error)
}

type Handlers struct {
	service Service

	log    *slog.Logger
	tracer trace.Tracer
}

func New(service Service, log *slog.Logger, tracer trace.Tracer) *Handlers {
	return &Handlers{
		log:     log,
		service: service,
		tracer:  tracer,
	}
}
