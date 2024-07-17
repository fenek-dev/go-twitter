package services

import (
	"context"

	"github.com/fenek-dev/go-twitter/src/common/models"
)

func (s *Services) FindTweetById(ctx context.Context, id string) (*models.Tweet, error) {
	ctx, span := s.tracer.Start(ctx, "tweet.service.FindTweetById")
	defer span.End()

	tweet, err := s.pg.FindTweetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (s *Services) CreateTweet(ctx context.Context, username, content string) (*models.Tweet, error) {
	ctx, span := s.tracer.Start(ctx, "tweet.service.CreateTweet")
	defer span.End()

	tweet, err := s.pg.CreateTweet(ctx, username, content)

	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (s *Services) UpdateTweet(ctx context.Context, id, content string) (*models.Tweet, error) {
	ctx, span := s.tracer.Start(ctx, "tweet.service.UpdateTweet")
	defer span.End()
	tweet, err := s.pg.UpdateTweet(ctx, id, content)

	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (s *Services) DeleteTweet(ctx context.Context, id string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "tweet.service.DeleteTweet")
	defer span.End()
	err := s.pg.DeleteTweet(ctx, id)

	if err != nil {
		return "", err
	}

	return id, nil
}
