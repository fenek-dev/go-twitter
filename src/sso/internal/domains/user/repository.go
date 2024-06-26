package user

import (
	"context"
	"fmt"
	"time"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) SaveUser(ctx context.Context, username string, passHash []byte) (models.User, error) {
	const op = "storage.pg.SaveUser"

	var user models.User

	rows, err := u.conn.Query(ctx, "INSERT INTO users(username, password, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING *",
		username,
		passHash,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (u *UserRepository) User(ctx context.Context, username string) (models.User, error) {
	const op = "storage.pg.User"

	var user models.User
	rows, _ := u.conn.Query(ctx, "SELECT * FROM users WHERE username = $1", username)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
