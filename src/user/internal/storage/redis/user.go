package redis

import (
	"context"
	"errors"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *Redis) User(ctx context.Context, id string) (*models.User, error) {
	ctx, span := r.tr.Start(ctx, "user.redis.User")
	defer span.End()

	key := "user." + id

	user := &models.User{}

	err := r.conn.HGetAll(ctx, key).Scan(user)
	if err != nil {
		span.RecordError(err)
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *Redis) SetUser(ctx context.Context, user *models.User, expiration time.Duration) error {
	ctx, span := r.tr.Start(ctx, "user.redis.SetUser")
	defer span.End()

	key := "user." + user.Id

	pipe := r.conn.Pipeline()
	pipe.HSet(ctx, key, user)
	pipe.Expire(ctx, key, expiration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
