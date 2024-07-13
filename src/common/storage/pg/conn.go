package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func New(ctx context.Context, DBUrl string) *pgx.Conn {

	conn, err := pgx.Connect(ctx, DBUrl)

	if err != nil {
		panic(fmt.Sprintf("can not connect to db: %s", err.Error()))
	}

	if err := conn.Ping(ctx); err != nil {
		panic(fmt.Sprintf("db ping didn't work: %s", err.Error()))
	}

	return conn
}
