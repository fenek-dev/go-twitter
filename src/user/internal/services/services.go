package services

import (
	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/user/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/user/internal/storage/redis"

	"go.opentelemetry.io/otel/trace"
)

type Services struct {
	sso ssov1.AuthServiceClient

	pg  *pg.Postgres
	rdb *redis.Redis

	tracer trace.Tracer
}

func New(sso ssov1.AuthServiceClient, pg *pg.Postgres, rdb *redis.Redis,
	tracer trace.Tracer) *Services {
	return &Services{
		sso:    sso,
		pg:     pg,
		rdb:    rdb,
		tracer: tracer,
	}
}
