package services

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common/models"
)

func (s *Services) FindUserById(ctx context.Context, id string) (*models.User, error) {
	ctx, span := s.tracer.Start(ctx, "tweet.service.FindTweetById")
	defer span.End()

	user, err := s.pg.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
