package storage

import (
	"context"
	"fmt"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) FindTweetById(ctx context.Context, id string) (models.Tweet, error) {
	const op = "read.tweet.findbyid"

	var tweet models.Tweet
	rows, err := s.conn.Query(ctx, "SELECT * FROM tweets WHERE id = $1", id)
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	tweet, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Tweet])
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}
