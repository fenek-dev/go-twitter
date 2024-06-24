package redis

import (
	"github.com/fenek-dev/go-twitter/src/cache/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	conn *redis.Client
}

func New(cfg *config.RedisConfig) *Redis {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &Redis{
		conn: conn,
	}
}

func (r *Redis) Close() error {
	return r.conn.Close()
}
