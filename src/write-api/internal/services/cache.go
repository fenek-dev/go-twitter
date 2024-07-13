package services

import (
	"context"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/mappers"
	"github.com/fenek-dev/go-twitter/src/common/models"
)

func (s *Services) CreateTweet(ctx context.Context, username, content string) (*models.Tweet, error) {
	ctx, span := s.tracer.Start(ctx, "write.service.CreateTweet")
	defer span.End()
	res, err := s.cache.CreateTweet(ctx, &proto.CreateTweetRequest{
		Username: username,
		Content:  content,
	})

	if err != nil {
		return nil, err
	}

	return mappers.ProtoTweetToModel(res.Tweet), nil
}

func (s *Services) UpdateTweet(ctx context.Context, id, content string) (*models.Tweet, error) {
	ctx, span := s.tracer.Start(ctx, "write.handler.UpdateTweet")
	defer span.End()
	res, err := s.cache.UpdateTweet(ctx, &proto.UpdateTweetRequest{
		Id:      id,
		Content: content,
	})

	if err != nil {
		return nil, err
	}

	return mappers.ProtoTweetToModel(res.Tweet), nil
}

func (s *Services) DeleteTweet(ctx context.Context, id string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "write.DeleteTweet")
	defer span.End()
	res, err := s.cache.DeleteTweet(ctx, &proto.DeleteTweetRequest{
		Id: id,
	})

	if err != nil {
		return "", err
	}

	return res.Id, nil
}
