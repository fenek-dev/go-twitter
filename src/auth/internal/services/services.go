package services

import (
	"log/slog"
	"time"

	"github.com/fenek-dev/go-twitter/src/auth/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/auth/internal/storage/redis"

	"go.opentelemetry.io/otel/trace"
)

type Services struct {
	pg  *pg.Postgres
	rdb *redis.Redis

	log      *slog.Logger
	tokenTTL time.Duration
	secret   string

	tracer trace.Tracer
}

func New(
	pg *pg.Postgres,
	rdb *redis.Redis,
	tracer trace.Tracer,
	log *slog.Logger,
	tokenTTL time.Duration,
	secret string,
) *Services {
	return &Services{
		pg:       pg,
		rdb:      rdb,
		tracer:   tracer,
		log:      log,
		tokenTTL: tokenTTL,
		secret:   secret,
	}
}
