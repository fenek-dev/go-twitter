package pg

import (
	"context"
	"fmt"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) FindUserById(ctx context.Context, id string) (models.User, error) {
	const op = "read.tweet.findbyid"

	var user models.User
	rows, err := p.conn.Query(ctx, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (p *Postgres) SaveUser(ctx context.Context, username string, passHash []byte) (string, error) {
	const op = "storage.pg.SaveUser"

	var usrname string
	err := p.conn.QueryRow(ctx, "INSERT INTO users(username, password) VALUES($1, $2) RETURNING username", username, passHash).Scan(&usrname)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return usrname, nil
}
