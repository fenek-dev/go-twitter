package pg

import (
	"context"
	"fmt"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) FindUserById(ctx context.Context, id string) (*models.User, error) {
	const op = "user.pg.FindUserById"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var user models.User
	queryCtx, span := p.tr.Start(ctx, op+".query")
	rows, err := p.conn.Query(queryCtx, "SELECT * FROM users WHERE id = $1", id)
	span.End()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, op+".scan")
	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	span.End()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
