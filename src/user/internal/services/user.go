package services

import (
	"context"
	"time"

	"github.com/fenek-dev/go-twitter/src/common/models"
)

func (s *Services) FindUserById(ctx context.Context, id string) (*models.User, error) {
	ctx, span := s.tracer.Start(ctx, "user.service.FindUserById")
	defer span.End()

	user, _ := s.rdb.User(ctx, id)
	if user != nil {
		return user, nil
	}

	user, err := s.pg.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.rdb.SetUser(ctx, user, time.Hour*24)
	if err != nil {
		span.RecordError(err)
	}

	return user, nil
}
