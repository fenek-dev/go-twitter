package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, DBUrl string, maxConns int32) *pgxpool.Pool {

	cfg, err := pgxpool.ParseConfig(DBUrl)
	if err != nil {
		panic(fmt.Sprintf("can not parse db config: %s", err.Error()))
	}
	cfg.MaxConns = maxConns

	conn, err := pgxpool.New(ctx, DBUrl)

	if err != nil {
		panic(fmt.Sprintf("can not connect to db: %s", err.Error()))
	}

	if err := conn.Ping(ctx); err != nil {
		panic(fmt.Sprintf("db ping didn't work: %s", err.Error()))
	}

	return conn
}
