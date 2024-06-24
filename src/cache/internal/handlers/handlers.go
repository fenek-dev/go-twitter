package handlers

import (
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/redis"
)

type Handlers struct {
	db    *pg.Postgres
	redis *redis.Redis
}

func New(db *pg.Postgres, redis *redis.Redis) *Handlers {
	return &Handlers{
		db:    db,
		redis: redis,
	}
}
