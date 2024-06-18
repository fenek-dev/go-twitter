package tweets

import (
	"context"
	"fmt"
	"time"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Create(ctx context.Context, username, content string) (models.Tweet, error) {
	const op = "write.tweet.create"

	var tweet models.Tweet
	rows, err := r.conn.Query(ctx, "INSERT INTO tweets(username, content, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING *",
		username,
		content,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	tweet, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Tweet])
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}

func (r *Repository) Update(ctx context.Context, id, content string) (models.Tweet, error) {
	const op = "write.tweet.update"

	var tweet models.Tweet
	rows, err := r.conn.Query(ctx, "UPDATE tweets SET content = $1, updated_at = $3 WHERE id = $2 RETURNING *",
		content,
		id,
		time.Now(),
	)
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	tweet, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Tweet])
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	const op = "write.tweet.delete"

	_, err := r.conn.Exec(ctx, "DELETE FROM tweets WHERE id = $1", id)

	return err
}
