package services

import (
	"context"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"go.opentelemetry.io/otel/trace"
)

type Postgres interface {
	FindTweetById(ctx context.Context, id string) (*models.Tweet, error)
	CreateTweet(ctx context.Context, username, content string) (*models.Tweet, error)
	UpdateTweet(ctx context.Context, id, content string) (*models.Tweet, error)
	DeleteTweet(ctx context.Context, id string) error
}

type Redis interface {
}

type Services struct {
	pg  Postgres
	rdb Redis

	tracer trace.Tracer
}

func New(pg Postgres, rdb Redis,
	tracer trace.Tracer) *Services {
	return &Services{
		pg:     pg,
		rdb:    rdb,
		tracer: tracer,
	}
}
