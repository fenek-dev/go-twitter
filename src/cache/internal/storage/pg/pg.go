package pg

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"
)

type Postgres struct {
	conn *pgx.Conn

	tr trace.Tracer
}

func New(ctx context.Context, url string, tr trace.Tracer) *Postgres {
	return &Postgres{
		conn: pg.New(ctx, url),
		tr:   tr,
	}
}

func (p *Postgres) Close(ctx context.Context) {
	p.conn.Close(ctx)
}
