package pg

import (
	"context"
	"fmt"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) FindUserById(ctx context.Context, id string) (models.User, error) {
	const op = "read.tweet.FindUserById"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var user models.User
	queryCtx, span := p.tr.Start(ctx, op+".query")
	rows, err := p.conn.Query(queryCtx, "SELECT * FROM users WHERE id = $1", id)
	span.End()
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, op+".scan")
	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	span.End()
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (p *Postgres) SaveUser(ctx context.Context, username string, passHash []byte) (string, error) {
	const op = "storage.pg.SaveUser"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var usrname string
	err := p.conn.QueryRow(ctx, "INSERT INTO users(username, password) VALUES($1, $2) RETURNING username", username, passHash).Scan(&usrname)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return usrname, nil
}
