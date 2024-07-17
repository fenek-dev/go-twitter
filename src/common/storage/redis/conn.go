package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	conn *redis.Client
}

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func New(ctx context.Context, cfg *Config) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	cmd := conn.Ping(ctx)
	if cmd.Err() != nil {
		panic(fmt.Sprintf("can not connect to redix: %s", cmd.Err().Error()))
	}

	return conn
}

func (r *Redis) Close() error {
	return r.conn.Close()
}
