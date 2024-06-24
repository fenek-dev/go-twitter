package storage

import (
	"context"
	"fmt"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) FindUserById(ctx context.Context, id string) (models.User, error) {
	const op = "read.tweet.findbyid"

	var user models.User
	rows, err := s.conn.Query(ctx, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
