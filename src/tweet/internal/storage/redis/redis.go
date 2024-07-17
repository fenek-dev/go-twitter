package redis

import (
	"context"

	rd "github.com/fenek-dev/go-twitter/src/common/storage/redis"
	"github.com/fenek-dev/go-twitter/src/tweet/config"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
)

type Redis struct {
	conn *redis.Client

	tr trace.Tracer
}

func New(ctx context.Context, cfg *config.RedisConfig, tr trace.Tracer) *Redis {
	return &Redis{
		conn: rd.New(ctx, &rd.Config{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
		tr: tr,
	}
}

func (p *Redis) Close() error {
	return p.conn.Close()
}
