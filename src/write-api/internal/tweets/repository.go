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

func (r *Repository) GetById() {

}

func (r *Repository) Create(ctx context.Context, username, content string) (*models.Tweet, error) {
	const op = "write.tweet.create"

	var tweet *models.Tweet
	err := r.conn.QueryRow(ctx, "INSERT INTO tweets(username, content, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING *",
		username,
		content,
		time.Now(),
		time.Now(),
	).Scan(&tweet)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}

func (r *Repository) Delete() {

}
