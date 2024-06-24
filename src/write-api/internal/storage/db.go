package storage

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn *pgx.Conn
}

func New(ctx context.Context, url string) *Storage {
	return &Storage{
		conn: pg.New(ctx, url),
	}
}

func (s *Storage) Close(ctx context.Context) {
	s.conn.Close(ctx)
}
