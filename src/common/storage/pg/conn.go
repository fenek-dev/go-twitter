package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func New(ctx context.Context, DBUrl string) *pgx.Conn {

	conn, err := pgx.Connect(ctx, DBUrl)

	if err != nil {
		panic("can not connect to db")
	}

	if err := conn.Ping(ctx); err != nil {
		panic("db ping didn't work")
	}

	return conn
}
