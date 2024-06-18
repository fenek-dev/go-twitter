package user

import (
	"context"
	"fmt"

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

func (u *UserRepository) SaveUser(ctx context.Context, username string, passHash []byte) (string, error) {
	const op = "storage.pg.SaveUser"

	var usrname string
	err := u.conn.QueryRow(ctx, "INSERT INTO users(username, password) VALUES($1, $2) RETURNING username", username, passHash).Scan(&usrname)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return usrname, nil
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
