package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) FindTweetById(ctx context.Context, id string) (models.Tweet, error) {
	const op = "read.tweet.FindTweetById"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var tweet models.Tweet

	queryCtx, span := p.tr.Start(ctx, op+".query")

	rows, err := p.conn.Query(queryCtx, "SELECT * FROM tweets WHERE id = $1", id)

	span.End()
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, op+".scan")

	tweet, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Tweet])

	span.End()
	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}

func (p *Postgres) CreateTweet(ctx context.Context, username, content string) (models.Tweet, error) {
	const op = "write.tweet.create"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var tweet models.Tweet
	queryCtx, span := p.tr.Start(ctx, op+".query")
	rows, err := p.conn.Query(queryCtx, "INSERT INTO tweets(username, content, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING *",
		username,
		content,
		time.Now(),
		time.Now(),
	)
	span.End()

	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, op+".scan")
	tweet, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Tweet])
	span.End()

	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}

func (p *Postgres) UpdateTweet(ctx context.Context, id, content string) (models.Tweet, error) {
	const op = "write.tweet.UpdateTweet"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	var tweet models.Tweet

	queryCtx, span := p.tr.Start(ctx, op+".query")
	rows, err := p.conn.Query(queryCtx, "UPDATE tweets SET content = $1, updated_at = $3 WHERE id = $2 RETURNING *",
		content,
		id,
		time.Now(),
	)
	span.End()

	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	_, span = p.tr.Start(ctx, op+".scan")
	tweet, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Tweet])
	span.End()

	if err != nil {
		return tweet, fmt.Errorf("%s: %w", op, err)
	}

	return tweet, nil
}

func (p *Postgres) DeleteTweet(ctx context.Context, id string) error {
	const op = "write.tweet.DeleteTweet"
	ctx, span := p.tr.Start(ctx, op)
	defer span.End()

	_, err := p.conn.Exec(ctx, "DELETE FROM tweets WHERE id = $1", id)

	return err
}
