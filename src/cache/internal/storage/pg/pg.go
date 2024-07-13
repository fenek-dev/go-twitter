package pg

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common/storage/pg"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/trace"
)

type Postgres struct {
	conn *pgxpool.Pool

	tr trace.Tracer
}

func New(ctx context.Context, url string, tr trace.Tracer) *Postgres {
	return &Postgres{
		conn: pg.New(ctx, url, 100),
		tr:   tr,
	}
}

func (p *Postgres) Close() {
	p.conn.Close()
}
