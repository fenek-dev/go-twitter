package pg

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn *pgx.Conn
}

func New(ctx context.Context, url string) *Postgres {
	return &Postgres{
		conn: pg.New(ctx, url),
	}
}

func (p *Postgres) Close(ctx context.Context) {
	p.conn.Close(ctx)
}
