package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(ctx context.Context, DBUrl string) *Storage {

	conn, err := pgx.Connect(ctx, DBUrl)

	if err != nil {
		panic("can not connect to db")
	}

	if err := conn.Ping(ctx); err != nil {
		panic("db ping didn't work")
	}

	return &Storage{db: conn}
}
