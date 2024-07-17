package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) SaveUser(ctx context.Context, username string, passHash []byte) (models.User, error) {
	const op = "storage.pg.SaveUser"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var user models.User

	queryCtx, span := p.tr.Start(ctx, "storage.pg.SaveUser.query")
	rows, err := p.conn.Query(queryCtx, "INSERT INTO users(username, password, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING *",
		username,
		passHash,
		time.Now(),
		time.Now(),
	)
	span.End()

	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, "storage.pg.User.scan")
	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	span.End()

	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (p *Postgres) User(ctx context.Context, username string) (models.User, error) {
	const op = "storage.pg.User"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var user models.User

	queryCtx, span := p.tr.Start(ctx, "storage.pg.User.query")
	rows, err := p.conn.Query(queryCtx, "SELECT * FROM users WHERE username = $1", username)
	span.End()

	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, "storage.pg.User.scan")
	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	span.End()
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
