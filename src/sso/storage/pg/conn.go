package pg

import (
	"context"
	"fmt"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(ctx context.Context, DBUrl string) *Storage {

	conn, err := pgx.Connect(ctx, DBUrl)

	if err != nil {
		panic("can not connect to db")
	}

	if err := conn.Ping(ctx); err != nil {
		panic("db ping didn't work")
	}

	return &Storage{db: conn}
}

func (s *Storage) SaveUser(ctx context.Context, username string, passHash []byte) (string, error) {
	const op = "storage.pg.SaveUser"

	var usrname string
	err := s.db.QueryRow(ctx, "INSERT INTO users(username, password) VALUES(?, ?) RETURNING username", username, passHash).Scan(&usrname)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return usrname, nil
}

func (s *Storage) User(ctx context.Context, username string) (models.User, error) {
	const op = "storage.pg.User"

	var user models.User
	err := s.db.QueryRow(ctx, "SELECT username, password FROM users WHERE username = ?", username).Scan(&user)

	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) Stop(ctx context.Context) {
	s.db.Close(ctx)
}
